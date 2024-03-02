package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	repo "github.com/satowo/todo-app/internal/repository"
)

func exitIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	exitIfError(err)

	// DBのセットアップ
	repo.SetUpDB()

	// エンドポイントとhandlerの登録
	// http.HandleFunc("/health", healthHandler)
	// http.HandleFunc("/allUsers", allUsersNameHandler)
	// http.HandleFunc("/channel", channelHandler)
	// http.HandleFunc("/message", MessageHandler)
	// http.HandleFunc("/messageEdit", MessageEditHandler)
	// http.HandleFunc("/messageDelete", MessageDeleteHandler)
	// http.HandleFunc("/channelMember", channelMemberHandler)
	// http.HandleFunc("/channelLeave", channelLeaveHandler)

	// HTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8080番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	err = http.ListenAndServe(":8083", nil)
	exitIfError(err)
}

// HTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		err := repo.DataBaseClose()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
