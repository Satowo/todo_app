package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	dmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func exitIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//DBのセットアップ
func SetUpDB() {
	// タイムゾーンの設定
	jst, err := time.LoadLocation(os.Getenv("TZ"))
	exitIfError(err)

	_db, err := openGORM(os.Getenv("DB_NAME"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "0.0.0.0", os.Getenv("MYSQL_PORT"), jst)
	exitIfError(err)

	db = _db
}

//GORMを使ってDBを開く
func openGORM(dbName, user, pass, host, port string, loc *time.Location) (*gorm.DB, error) {
	m := mysql.New(mysql.Config{
		DSNConfig: &dmysql.Config{
			User:      user,
			Passwd:    pass,
			DBName:    dbName,
			Addr:      fmt.Sprintf("%s:%s", host, port),
			Net:       "tcp",
			Loc:       loc,
			ParseTime: true,
		},
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportNullAsDefaultValue: true,
		DontSupportRenameColumnUnique: true,
	})

	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

//DBを閉じる
func DataBaseClose() error {
	sqlDB, err := db.DB()
    sqlDB.Close()
	return err
}
