package models

import (
	"fmt"
	"github.com/GaoHaHa-IronMan/go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int   `gorm:"primary_key" json:"id"`
	CreatedOn  int64 `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn int64 `gorm:"autoUpdateTime" json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   tablePrefix, // 表名前缀，`User`表为`t_users`
					SingularTable: true,        // 使用单数表名，启用该选项后，`User` 表将是`user`
				},
			})
	}

	if err != nil {
		log.Fatal(err)
	}

	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

//func updateTimeStampForCreateCallback(db *gorm.DB) {
//	if db.Statement.Schema != nil {
//		nowTime := time.Now().Unix()
//		if createTimeField, ok := db.Statement.Schema.FieldsByName["CreatedOn"]; ok {
//
//			err := createTimeField.Set(db.Statement.ReflectValue, nowTime)
//			if err != nil {
//				logging.Error(err)
//			}
//		}
//		if createTimeField, ok := db.Statement.Schema.FieldsByName["ModifiedOn"]; ok {
//
//			err := createTimeField.Set(db.Statement.ReflectValue, nowTime)
//			if err != nil {
//				logging.Error(err)
//			}
//		}
//	}
//}
//
//func updateTimeStampForUpdateCallback(db *gorm.DB) {
//
//	if _, ok := db.Get("gorm:update_column"); !ok {
//		db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
//	}
//}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
}
