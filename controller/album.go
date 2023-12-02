package controller

import (
	"encoding/json"
	//"fmt"
	"net/http"
	//"strconv"

	//"github.com/gorilla/mux"
	//"server-recruit-challenge-sample/model"
	"server-recruit-challenge-sample/service"
)

// albumController 構造体は、service.AlumService インターフェースを持ち、アルバムに関するHTTPリクエストを処理
type albumController struct {
	service service.AlbumService
}

// NewSingerController 関数：singerController インスタンスを作成して返す
func NewAlbumController(s service.AlbumService) *albumController {
	return &albumController{service: s}
}

// GET /albums のハンドラー
// GETリクエストを処理してアルバムリストを取得し、JSON形式でレスポンスを返す
func (c *albumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.service.GetAlbumListService(r.Context()) // service/album.go ファイルの GetAlbumListService メソッドを呼び出す
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(albums)
}

// // GET /singers/{id} のハンドラー
// // GETリクエストを処理して歌手を取得し、JSON形式でレスポンスを返す
// func (c *singerController) GetSingerDetailHandler(w http.ResponseWriter, r *http.Request) {
// 	singerID, err := strconv.Atoi(mux.Vars(r)["id"]) // URLパラメータから歌手IDを取得
// 	if err != nil {
// 		err = fmt.Errorf("invalid path param: %w", err)
// 		errorHandler(w, r, 400, err.Error())
// 		return
// 	}

// 	// service/singer.go ファイルの GetSingerService メソッドを呼び出す
// 	singer, err := c.service.GetSingerService(r.Context(), model.SingerID(singerID))
// 	if err != nil {
// 		errorHandler(w, r, 500, err.Error())
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(singer)
// }

// // POST /singers のハンドラー
// // POSTリクエストを処理して歌手を登録し、JSON形式でレスポンスを返す
// func (c *singerController) PostSingerHandler(w http.ResponseWriter, r *http.Request) {
// 	var singer *model.Singer
// 	if err := json.NewDecoder(r.Body).Decode(&singer); err != nil { // リクエストボディから歌手データを取得
// 		err = fmt.Errorf("invalid body param: %w", err) // リクエストボディが不正な場合はエラーを返す
// 		errorHandler(w, r, 400, err.Error())
// 		return
// 	}

// 	if err := c.service.PostSingerService(r.Context(), singer); err != nil { // service/singer.go ファイルの PostSingerService メソッドを呼び出す
// 		errorHandler(w, r, 500, err.Error())
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(singer)
// }

// // DELETE /singers/{id} のハンドラー
// // DELETEリクエストを処理して歌手を削除する
// func (c *singerController) DeleteSingerHandler(w http.ResponseWriter, r *http.Request) {
// 	singerID, err := strconv.Atoi(mux.Vars(r)["id"]) // URLパラメータから歌手IDを取得
// 	if err != nil {
// 		err = fmt.Errorf("invalid path param: %w", err)
// 		errorHandler(w, r, 400, err.Error())
// 		return
// 	}

// 	// service/singer.go ファイルの DeleteSingerService メソッドを呼び出す
// 	if err := c.service.DeleteSingerService(r.Context(), model.SingerID(singerID)); err != nil {
// 		errorHandler(w, r, 500, err.Error())
// 		return
// 	}
// 	w.WriteHeader(204)
// }
