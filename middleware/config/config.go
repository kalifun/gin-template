package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kalifun/gin-template/global"
	"github.com/kalifun/gin-template/utils"
	"github.com/spf13/viper"
)

const ConfigPath = "config.yaml"

func Init() {
	v := viper.New()
	v.SetConfigFile(ConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监听配置文件的修改和变动
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 当出现配置改动时 重新unmarshal到struct对象
		if err := v.Unmarshal(&global.ConfigSvr); err != nil {
			fmt.Println(err)
		}
	})
	//将配置解析为Struct对象
	if err := v.Unmarshal(&global.ConfigSvr); err != nil {
		fmt.Println(err)
	}
	global.Viper = v
	printStr := `Init Config ` + utils.Green("Success")
	fmt.Println(printStr)
}
