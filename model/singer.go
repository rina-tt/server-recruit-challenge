// 歌手（Singer）に関するデータモデルを定義するためのパッケージ

package model // このファイルが model パッケージであることを示す

type SingerID int // 歌手（Singer）の ID

type Singer struct { // 歌手（Singer）の構造体
	ID   SingerID `json:"id"`
	Name string   `json:"name"`
}