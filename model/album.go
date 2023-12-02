// アルバム（Album）に関するデータモデルを定義するためのパッケージ

package model // このファイルが model パッケージであることを示す

type AlbumID int // アルバム（Album）の ID

type Album struct { // アルバム（Album）の構造体
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	SingerID SingerID `json:"singer_id"` // モデル Singer の ID と紐づきます
}
