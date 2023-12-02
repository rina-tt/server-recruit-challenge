// HTTPリクエストのルーティングを設定するためのルーターを定義

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/memorydb"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

// 新しい mux.Router インスタンスを作成し、それに対して歌手に関するエンドポイントのハンドラーを設定
func NewRouter() *mux.Router {
	singerRepo := memorydb.NewSingerRepository() // infra/memorydb/singer.go ファイルの NewSingerRepository 関数を呼び出す
	singerService := service.NewSingerService(singerRepo) // service/singer.go ファイルの NewSingerService 関数を呼び出す
	singerController := controller.NewSingerController(singerService) // controller/singer.go ファイルの NewSingerController 関数を呼び出す

	r := mux.NewRouter()

	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet) // GET /singers のハンドラー
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet) // GET /singers/{id} のハンドラー
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost) // POST /singers のハンドラー
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete) // DELETE /singers/{id} のハンドラー

	r.Use(middleware.LoggingMiddleware) // ログ出力用のミドルウェアを適用

	return r
}
