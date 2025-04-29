package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gintemplate/app/config"
)

var DB *gorm.DB

func InitMySQL() error {
	mysqlConfig := config.ConfigInfo.DB.MySQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	)
	//fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	err = DB.AutoMigrate()
	if err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	return nil
}
