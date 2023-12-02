// HTTPリクエストでエラーが発生した場合の共通的なエラーハンドリングを提供するためのファイル

package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// エラーが発生したときのレスポンス処理をここで行う
// w http.ResponseWriter：HTTPレスポンスを書き込むための構造体
// r *http.Request：HTTPリクエストを表す構造体
// statusCode int：HTTPステータスコード
// message string：エラーメッセージ
func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Printf("error: %s\n", message) // エラーをログに出力する

	type ErrorMessage struct { // エラーメッセージをJSON形式で返す
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorMessage{Message: message})
}
