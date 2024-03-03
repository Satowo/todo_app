package main

import (
	"github.com/satowo/todo-app/internal/server"
)

func main() {
	config := server.NewConfig()

	httpServer := server.NewServer(config)

	httpServer.Run()
}


// // 環境変数の読み込み
// dbConfig := setup.GetDBEnv()

// // DBのセットアップ
// err := setup.SetUpDB(dbConfig)
// exitIfError(err)

// // サーバーのセットアップ
// setup.NewServer()

// // HTTPサーバー停止時にDBをクローズする
// setup.CloseDBWithSysCall()

// // 8080番ポートでリクエストを待ち受ける
// log.Println("Server is running on ")

// err = http.ListenAndServe(":8083", nil)
// exitIfError(err)