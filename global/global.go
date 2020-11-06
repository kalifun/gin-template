package global

import (
	"github.com/kalifun/gin-template/config"
	oplog "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	Viper     *viper.Viper
	ConfigSvr config.ConfigInfo
	Log        *oplog.Logger
)
