package dao

import (
	"database/sql"

	"github.com/starichat/jianshu/global"
)

var masterdb *sql.DB

func init() {
	// 保证viper配置文件已处理
	global.Init()
}
