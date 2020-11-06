package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kalifun/gin-template/global"
)

/**
 * MySQL 配置
 * 注册驱动
 * orm.RegisterDriver("mysql", orm.DR_MySQL)
 * mysql用户：root ，root的秘密：tom ， 数据库名称：test ， 数据库别名：default
 * orm.RegisterDataBase("default", "mysql", "root:tom@/test?charset=utf8")
 */
func Init() {
	sqlSet := global.ConfigSvr.Mysql
	sqlConn := sqlSet.UserName + ":" + sqlSet.PassWord + "@(" + sqlSet.Path + ")/" + sqlSet.DbName + "?" + sqlSet.Config
	db, err := gorm.Open("mysql", sqlConn)
	if err != nil {
		global.Log.Error(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
