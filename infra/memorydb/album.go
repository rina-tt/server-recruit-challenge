// メモリ内でアルバムデータを保持するためのデータベース（インメモリデータベース）を実装するためのパッケージ

package memorydb

import (
	"context"
	//"errors"
	"sync"

	"server-recruit-challenge-sample/model"
	"server-recruit-challenge-sample/repository"
)

// sync.RWMutex を埋め込み、albumMap フィールドでアルバムデータを保持
// AlbumID をキーとし、model.Album を値とするマップ
type albumRepository struct {
	sync.RWMutex
	albumMap map[model.AlbumID]*model.Album // キーが AlbumID、値が model.Album のマップ
}

// インターフェースが正しく実装されていることを確認するためのコード
var _ repository.AlbumRepository = (*albumRepository)(nil)

// 初期化済みのアルバムデータを持つ albumRepository インスタンスを返す
func NewAlbumRepository() *albumRepository {
	var initMap = map[model.AlbumID]*model.Album{
		1: {ID: 1, Title: "Alice's 1st Album", SingerID: 1},
		2: {ID: 2, Title: "Alice's 2nd Album", SingerID: 1},
		3: {ID: 3, Title: "Bella's 1st Album", SingerID: 2},
	}

	return &albumRepository{
		albumMap: initMap,
	}
}

// GetAll はアルバムデータを全件取得する。読み取り用のロックを取得し、アルバムデータをスライスにコピーして返す。
func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
	r.RLock()
	defer r.RUnlock()

	albums := make([]*model.Album, 0, len(r.albumMap))
	for _, s := range r.albumMap {
		albums = append(albums, s)
	}
	return albums, nil
}

// // Get は歌手ADに対応する歌手データを取得する。読み取り用のロックを取得し、指定されたIDの歌手が存在しない場合はエラーを返す。
// func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
// 	r.RLock()
// 	defer r.RUnlock()

// 	singer, ok := r.singerMap[id]
// 	if !ok {
// 		return nil, errors.New("not found")
// 	}
// 	return singer, nil
// }

// // Add は新しい歌手を追加する。書き込み用のロックを取得し、歌手を singerMap に追加する。
// func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
// 	r.Lock()
// 	r.singerMap[singer.ID] = singer
// 	r.Unlock()
// 	return nil
// }

// // Delete は指定された歌手IDに対応する歌手を削除する。書き込み用のロックを取得し、singerMap から指定されたIDの歌手を削除する
// func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
// 	r.Lock()
// 	delete(r.singerMap, id)
// 	r.Unlock()
// 	return nil
// }
