package setup

import (
	"fmt"
	"time"

	dmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//GORMを使ってDBを開く
func SetUpDB(dbConfig *DBConfig) error {
	m := mysql.New(mysql.Config{
		DSNConfig: &dmysql.Config{
			User:      dbConfig.DBUser,
			Passwd:    dbConfig.DBPass,
			DBName:    dbConfig.DBName,
			Addr:      fmt.Sprintf("%s:%s", dbConfig.DBHost, dbConfig.DBPort),
			Net:       "tcp",
			Loc:       time.Local,
			ParseTime: true,
		},
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportNullAsDefaultValue: true,
		DontSupportRenameColumnUnique: true,
	})

	_db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	db = _db
	return nil
}

//DBを閉じる
func DataBaseClose() error {
	sqlDB, err := db.DB()
    sqlDB.Close()
	return err
}
