package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"server-recruit-challenge-sample/model"
	"server-recruit-challenge-sample/service"
)

// albumController 構造体は、service.AlumService インターフェースを持ち、アルバムに関するHTTPリクエストを処理
type albumController struct {
	service service.AlbumService
}

// NewAlbumController 関数：albumController インスタンスを作成して返す
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

// GET /albums/{id} のハンドラー
// GETリクエストを処理してアルバムを取得し、JSON形式でレスポンスを返す
func (c *albumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
	albumID, err := strconv.Atoi(mux.Vars(r)["id"]) // URLパラメータからアルバムIDを取得
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}


	// service/album.go ファイルの GetAlbumService メソッドを呼び出す
	album, err := c.service.GetAlbumService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// POST /albums のハンドラー
// POSTリクエストを処理してアルバムを登録し、JSON形式でレスポンスを返す
func (c *albumController) PostAlbumHandler(w http.ResponseWriter, r *http.Request) {
	var album *model.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil { // リクエストボディからアルバムデータを取得
		err = fmt.Errorf("invalid body param: %w", err) // リクエストボディが不正な場合はエラーを返す
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.PostAlbumService(r.Context(), album); err != nil { // service/album.go ファイルの PostAlbumService メソッドを呼び出す
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// DELETE /albums/{id} のハンドラー
// DELETEリクエストを処理してアルバムを削除する
func (c *albumController) DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	albumID, err := strconv.Atoi(mux.Vars(r)["id"]) // URLパラメータから歌手IDを取得
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	// service/album.go ファイルの DeleteAlbumService メソッドを呼び出す
	if err := c.service.DeleteAlbumService(r.Context(), model.AlbumID(albumID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.WriteHeader(204)
}
