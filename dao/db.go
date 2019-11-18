package dao

import "database/sql"

var masterdb *sql.DB


fun init() {
	// 保证viper配置文件已处理
	global.Init()
}