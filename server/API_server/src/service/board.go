package service

import (
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type BoardService struct {
	BoardRepository repository.BoardRepository
}

func NewBoardService(repository repository.BoardRepository) BoardService {
	return BoardService{BoardRepository: repository}
}

func (boardSvc *BoardService) GetBoard(id string) []model.Board {
	return boardSvc.BoardRepository.SelectByUserId(id)
}

func (boardSvc *BoardService) CreateNewBoard(userId, title, img string) error {
	// ボード新規作成機能
	err := boardSvc.BoardRepository.InsertBoard(userId, title, img)
	if err != nil {
		return err
	}

	return nil
}

func (boardSvc *BoardService) ChangePublishInfo(userId string, boardId int, publish bool) error {
	// ボードの公開・非公開情報更新機能
	return boardSvc.BoardRepository.UpdateBoardPublish(userId, boardId, publish)
}

func (boardSvc *BoardService) SendInviteMail(boardId int, token, email, message string) error {
	//var userId int//
	// 受け取ったTokenを元にuserIdを取得し、ボードへのアクセス権限があるかを調べる

	// 招待メール送信機能
	return nil
}
