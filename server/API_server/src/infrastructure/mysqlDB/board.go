package mysqlDB

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type BoardRepositoryImpliment struct {
	DB *gorm.DB
}

func NewBoardRepoImpl(DB *gorm.DB) repository.BoardRepository {
	return &BoardRepositoryImpliment{
		DB: DB,
	}
}

func (p *BoardRepositoryImpliment) SelectByUserId(userId string) []model.Board {
	//ユーザIDにひもづくボードを全件取得する

	//モデル変数の宣言
	var boards []model.Board
	var user model.User

	//userIdによるユーザー情報の取得
	p.DB.Where("id=?", userId).Find(&user)

	//取得したユーザーに対応するBoardを取得
	p.DB.Model(&user).Related(&boards, "Boards")

	return boards
}

func (p *BoardRepositoryImpliment) InsertBoard(userId int, title, img string) error {
	// ボード新規作成
	return nil
}

func (p *BoardRepositoryImpliment) UpdateBoardPublish(userId string, boardId int, publish bool) error {
	// ボードの公開・非公開情報更新

	//モデル変数の宣言
	var user model.User
	var boards []model.Board

	//userIdによるユーザー情報の取得
	p.DB.Where("id=?", userId).Find(&user)

	//取得したユーザーに対応するBoardを取得
	p.DB.Model(&user).Related(&boards, "Boards")

	for _, board := range boards {
		//該当のBoardを見つけたらUPDATEをかける
		if board.Id == boardId {
			p.DB.Model(&board).Update("publish", publish)
		}
	}

	return nil
}
