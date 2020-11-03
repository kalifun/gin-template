package global

import (
	"github.com/kalifun/gin-template/config"
	"github.com/spf13/viper"
)

var (
	Viper     *viper.Viper
	ConfigSvr config.ConfigInfo
)
