package logs

import (
	"fmt"
	"github.com/kalifun/gin-template/config"
	"github.com/kalifun/gin-template/global"
	"github.com/kalifun/gin-template/utils"
	rotates "github.com/lestrrat-go/file-rotatelogs"
	oplog "github.com/op/go-logging"
	"io"
	"os"
	"strings"
	"time"
)

const (
	logDir      = "logs"
	logSoftLink = "latest_log"
	module      = "auth"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold} ▶ [%{level:.6s}] %{message}%{color:reset}`
)

func InitLog() {
	c := global.ConfigSvr.Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := oplog.MustGetLogger(module)
	var backends []oplog.Backend
	registerStdout(c, &backends)
	if fileWriter := registerFile(c, &backends); fileWriter != nil {
		io.MultiWriter(fileWriter, os.Stdout)
	}
	oplog.SetBackend(backends...)
	global.Log = logger
}

func registerStdout(c config.Log, backends *[]oplog.Backend) {
	if c.Stdout != "" {
		level, err := oplog.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		*backends = append(*backends, createBackend(os.Stdout, c, level))
	}
}

func registerFile(c config.Log, backends *[]oplog.Backend) io.Writer {
	if c.File != "" {
		if ok, _ := utils.PathExists(logDir); !ok {
			// directory not exist
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		fileWriter, err := rotates.New(
			logDir+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
			// generate soft link, point to latest log file
			rotates.WithLinkName(logSoftLink),
			// maximum time to save log files
			rotates.WithMaxAge(7*24*time.Hour),
			// time period of log file switching
			rotates.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
		}
		level, err := oplog.LogLevel(c.File)
		if err != nil {
			fmt.Println(err)
		}
		*backends = append(*backends, createBackend(fileWriter, c, level))

		return fileWriter
	}
	return nil
}

func createBackend(w io.Writer, c config.Log, level oplog.Level) oplog.Backend {
	backend := oplog.NewLogBackend(w, c.Prefix, 0)
	stdoutWriter := false
	if w == os.Stdout {
		stdoutWriter = true
	}
	format := getLogFormatter(c, stdoutWriter)
	backendLeveled := oplog.AddModuleLevel(oplog.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled
}

func getLogFormatter(c config.Log, stdoutWriter bool) oplog.Formatter {
	pattern := defaultFormatter
	if !stdoutWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		// Remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return oplog.MustStringFormatter(pattern)
}

/*
写gin的日志文件
*/
//func WriteGinLog() gin.HandlerFunc {
//
//}
