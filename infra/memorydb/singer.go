// メモリ内で歌手データを保持するための簡単なデータベース（インメモリデータベース）を実装するパッケージ

package memorydb // このファイルが memorydb パッケージであることを示す

import (
	"context"
	"errors"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

// singerRepository 構造体は：sync.RWMutex を埋め込み、singerMap フィールドで歌手データを保持
// SingerID をキーとし、model.Singer を値とするマップ
type singerRepository struct {
	sync.RWMutex
	singerMap map[model.SingerID]*model.Singer // キーが SingerID、値が model.Singer のマップ
}

// インターフェースが正しく実装されていることを確認するためのコード
var _ repository.SingerRepository = (*singerRepository)(nil)

// 歌手データを保持するための簡単なデータベース（インメモリデータベース）を初期化する
func NewSingerRepository() *singerRepository {
	var initMap = map[model.SingerID]*model.Singer{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bella"},
		3: {ID: 3, Name: "Chris"},
		4: {ID: 4, Name: "Daisy"},
		5: {ID: 5, Name: "Ellen"},
	}

	return &singerRepository{
		singerMap: initMap,
	}
}

// GetAll は歌手データを全件取得する。読み取り用のロックを取得し、歌手データをスライスにコピーして返す。
func (r *singerRepository) GetAll(ctx context.Context) ([]*model.Singer, error) {
	r.RLock()
	defer r.RUnlock()

	singers := make([]*model.Singer, 0, len(r.singerMap))
	for _, s := range r.singerMap {
		singers = append(singers, s)
	}
	return singers, nil
}

// Get は歌手ADに対応する歌手データを取得する。読み取り用のロックを取得し、指定されたIDの歌手が存在しない場合はエラーを返す。
func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	r.RLock()
	defer r.RUnlock()

	singer, ok := r.singerMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return singer, nil
}

// Add は新しい歌手を追加する。書き込み用のロックを取得し、歌手を singerMap に追加する。
func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
	r.Lock()
	r.singerMap[singer.ID] = singer
	r.Unlock()
	return nil
}

// Delete は指定された歌手IDに対応する歌手を削除する。書き込み用のロックを取得し、singerMap から指定されたIDの歌手を削除する
func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
	r.Lock()
	delete(r.singerMap, id)
	r.Unlock()
	return nil
}
