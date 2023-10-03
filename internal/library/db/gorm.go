package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type Config struct {
	Host          string `help:"数据库地址" default:"localhost"`
	Port          int    `help:"数据库端口" devDefault:"3306" testDefault:"3307" releaseDefault:"3306"`
	Username      string `help:"数据库帐号" default:"root"`
	Password      string `help:"数据库密码" default:"root"`
	Database      string `help:"数据库名称" default:"cms"`
	TablePrefix   string `help:"数据库表前缀" default:""`
	Charset       string `help:"数据库编码" default:"utf8mb4"`
	SingularTable bool   `help:"是否使用单数表名" default:"true"`
	TimeZone      string `help:"时区" default:"Asia/Shanghai"`
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset, true)
}

func NewGormDB(conf *Config) (DB *gorm.DB, err error) {
	DB, err = gorm.Open(mysql.Open(conf.Dsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix,   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: conf.SingularTable, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
		Logger: logger.Default.LogMode(logger.Info), // 打印日志，不需要注释
	})
	if err != nil {
		fmt.Println("打开mysql失败", err)
		return
	}

	//db.SingularTable(true)
	sqlDB, err := DB.DB()
	if err != nil || sqlDB == nil {
		fmt.Println("打开db失败", err)
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
}
