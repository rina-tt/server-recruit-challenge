// 歌手（Singer）に関するデータの永続化と取得のためのリポジトリ（Repository）を定義するパッケージ
// データベースや他の永続化メカニズムに依存せず、歌手データに対する基本的な操作を抽象化

package repository // このファイルが repository パッケージであることを示す

import (
	"context"

	"server-recruit-challenge-sample/model"
)

// SingerRepository インターフェース：歌手に関するデータの永続化と取得に必要な基本的なメソッドを定義
type SingerRepository interface {
	GetAll(ctx context.Context) ([]*model.Singer, error)               // すべての歌手を取得
	Get(ctx context.Context, id model.SingerID) (*model.Singer, error) // 指定された歌手IDに対応する歌手を取得
	Add(ctx context.Context, singer *model.Singer) error               // 新しい歌手を追加
	Delete(ctx context.Context, id model.SingerID) error               // 指定された歌手IDに対応する歌手を削除
}
