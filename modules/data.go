package modules

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Engine TODO
var Engine *gorm.DB

// Connection TODO
func Connection(user, pass, host, name, char string, port int) (*gorm.DB, error) {
	slowLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 1 * time.Microsecond,
			// 设置日志级别，只有 Warn 和 Info 级别会输出慢查询日志
			LogLevel: logger.Error,
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		name,
		char,
	)
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: slowLogger,
	})
	if err != nil {
		fmt.Errorf("mysql connect fail address is %s", dsn)
	}

	sqlDB, _ := engine.DB()
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetMaxIdleConns(20)
	return engine, err
}

type Data struct {
	ID         uint64    `json:"id,omitempty"        gorm:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	AnimalType string    `json:"animal_type" gorm:"column:animal_type;not null"`
	Eat        string    `json:"eat" gorm:"column:eat;not null"`
	Move       string    `json:"move" gorm:"column:move;not null"`
	Speak      string    `json:"speak" gorm:"column:speak;not null"`
}

// 定义模型的数据表名称
func (data *Data) TableName() string {
	return "data"
}
