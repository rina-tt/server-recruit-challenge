// 課題3 新規作成
// アルバム（Album）に関するデータの永続化と取得のためのリポジトリ（Repository）を定義するパッケージ

package repository // このファイルが repository パッケージであることを示す

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

// AlbumRepository インターフェース：アルバムに関するデータの永続化と取得に必要な基本的なメソッドを定義
type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.Album, error)               // すべての歌手を取得
	Get(ctx context.Context, id model.AlbumID) (*model.Album, error) // 指定されたアルバムIDに対応する歌手を取得
	//Add(ctx context.Context, singer *model.Album) error               // 新しいアルバムを追加
	//Delete(ctx context.Context, id model.AlbumID) error               // 指定されたアルバムIDに対応する歌手を削除
}
