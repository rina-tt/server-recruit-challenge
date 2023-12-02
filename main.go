package main // このファイルが main パッケージであることを示

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pulse227/server-recruit-challenge-sample/api"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// api パッケージ内の NewRouter 関数を呼び出して、新しいルーターを作成
	r := api.NewRouter()

	// HTTPサーバーの設定
	server := &http.Server{
		Addr:    ":8888", // ポート番号を指定
		Handler: r,       // ルーターをハンドラーとして設定
	}
	// ゴルーチン（非同期処理）を開始し、割り込み（os.Interrupt）が発生した場合にサーバーを graceful にシャットダウン
	go func() {
		<-ctx.Done() // 割り込みが発生するまで待機
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5秒間のタイムアウトを設定
		defer cancel()
		server.Shutdown(ctx) // シャットダウン
	}()
	log.Println("server start running at :8888") // ログを出力
	log.Fatal(server.ListenAndServe()) // サーバーを起動 (エラーが発生した場合はログを出力して終了)
}
