package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/satowo/todo-app/internal/setup"
)

func exitIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// 環境変数の読み込み
	dbConfig := setup.GetDBEnv()

	// DBのセットアップ
	err := setup.SetUpDB(dbConfig)
	exitIfError(err)

	// サーバーのセットアップ
	setup.NewServer()

	// HTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8080番ポートでリクエストを待ち受ける
	log.Println("Server is running on :8083")

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

		err := setup.DataBaseClose()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
