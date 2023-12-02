// 歌手（Singer）に関するサービスを提供するためのパッケージ

package service // このファイルが service パッケージであることを示す

import (
	"context"

	"server-recruit-challenge-sample/model"
	"server-recruit-challenge-sample/repository"
)


// SingerService は歌手（Singer）に関するサービスを提供するためのインターフェース
type SingerService interface {
	GetSingerListService(ctx context.Context) ([]*model.Singer, error) // 一覧を取得する
	GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error) // 取得する
	PostSingerService(ctx context.Context, singer *model.Singer) error // 追加する
	DeleteSingerService(ctx context.Context, singerID model.SingerID) error // 削除する
}


// 歌手（Singer）に関するサービスを提供するための構造体
type singerService struct {
	// repository/singer.go ファイルの SingerRepository インターフェースを埋め込む
	singerRepository repository.SingerRepository
}


// 構造体 singerService が SingerService インターフェースを実装していることをコンパイラに伝える
var _ SingerService = (*singerService)(nil)


// NewSingerService は歌手（Singer）に関するサービスを提供するための構造体を生成する
func NewSingerService(singerRepository repository.SingerRepository) *singerService {
	return &singerService{singerRepository: singerRepository}
}


// 以下、サービスメソッドの実装

// 歌手（Singer）の一覧を取得するサービスメソッド
func (s *singerService) GetSingerListService(ctx context.Context) ([]*model.Singer, error) {
	singers, err := s.singerRepository.GetAll(ctx) // repository/singer.go ファイルの GetAll メソッドを呼び出す
	if err != nil {
		return nil, err
	}
	return singers, nil
}


// 指定された歌手IDに対応する歌手（Singer）を取得するサービスメソッド
func (s *singerService) GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error) {
	singer, err := s.singerRepository.Get(ctx, singerID) // repository/singer.go ファイルの Get メソッドを呼び出す
	if err != nil {
		return nil, err
	}
	return singer, nil
}


// 新しい歌手（Singer）を追加するサービスメソッド
func (s *singerService) PostSingerService(ctx context.Context, singer *model.Singer) error {
	if err := s.singerRepository.Add(ctx, singer); err != nil { // repository/singer.go ファイルの Add メソッドを呼び出す
		return err
	}
	return nil
}


// 指定された歌手IDに対応する歌手（Singer）を削除するサービスメソッド
func (s *singerService) DeleteSingerService(ctx context.Context, singerID model.SingerID) error {
	if err := s.singerRepository.Delete(ctx, singerID); err != nil { // repository/singer.go ファイルの Delete メソッドを呼び出す
		return err
	}
	return nil
}
