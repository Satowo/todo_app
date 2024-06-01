package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	dmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//GORMを使ってDBを開く
func SetUpDB() (*gorm.DB, error) {
	m := mysql.New(mysql.Config{
		DSNConfig: &dmysql.Config{
			User:      config.DBConfig.DBUser,
			Passwd:    config.DBConfig.DBPass,
			DBName:    config.DBConfig.DBName,
			Addr:      fmt.Sprintf("%s:%s", config.DBConfig.DBHost, config.DBConfig.DBPort),
			Net:       "tcp",
			Loc:       time.Local,
			ParseTime: true,
		},
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportNullAsDefaultValue: true,
		DontSupportRenameColumnUnique: true,
	})

	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	
	return db, nil
}

// HTTPサーバー停止時にDBを閉じる
func CloseDBWithSysCall(db *gorm.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		err := DataBaseClose(db)
		if err != nil {
			log.Fatal(err)
		}
	
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

func DataBaseClose(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}