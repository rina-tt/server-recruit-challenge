// アルバム（Album）に関するサービスを提供するためのパッケージ

package service // このファイルが service パッケージであることを示す

import (
	"context"

	"server-recruit-challenge-sample/model"
	"server-recruit-challenge-sample/repository"
)


// AlbumService はアルバム（Album）に関するサービスを提供するためのインターフェース
type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error) // 一覧を取得する
	//GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Singer, error) // 取得する
	//PostAlbumService(ctx context.Context, album *model.Album) error // 追加する
	//DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error // 削除する
}


// アルバム（Album）に関するサービスを提供するための構造体
type albumService struct {
	// repository/album.go ファイルの AlbumRepository インターフェースを埋め込む
	albumRepository repository.AlbumRepository
}


// 構造体 albumService が AlbumService インターフェースを実装していることをコンパイラに伝える
var _ AlbumService = (*albumService)(nil)


// NewAlbumService はアルバム（Album）に関するサービスを提供するための構造体を生成する
func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{albumRepository: albumRepository}
}


// 以下、サービスメソッドの実装

// アルバム（Album）の一覧を取得するサービスメソッド
func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albums, err := s.albumRepository.GetAll(ctx) // repository/album.go ファイルの GetAll メソッドを呼び出す
	if err != nil {
		return nil, err
	}
	return albums, nil
}


// // 指定された歌手IDに対応する歌手（Singer）を取得するサービスメソッド
// func (s *singerService) GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error) {
// 	singer, err := s.singerRepository.Get(ctx, singerID) // repository/singer.go ファイルの Get メソッドを呼び出す
// 	if err != nil {
// 		return nil, err
// 	}
// 	return singer, nil
// }


// // 新しい歌手（Singer）を追加するサービスメソッド
// func (s *singerService) PostSingerService(ctx context.Context, singer *model.Singer) error {
// 	if err := s.singerRepository.Add(ctx, singer); err != nil { // repository/singer.go ファイルの Add メソッドを呼び出す
// 		return err
// 	}
// 	return nil
// }


// // 指定された歌手IDに対応する歌手（Singer）を削除するサービスメソッド
// func (s *singerService) DeleteSingerService(ctx context.Context, singerID model.SingerID) error {
// 	if err := s.singerRepository.Delete(ctx, singerID); err != nil { // repository/singer.go ファイルの Delete メソッドを呼び出す
// 		return err
// 	}
// 	return nil
// }
