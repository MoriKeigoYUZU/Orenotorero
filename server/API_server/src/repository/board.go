package repository

import (
	"orenotorero/db/Model"
)

type BoardRepository interface {
	SelectByUserId(userId string) []model.Board
	InsertBoard(userId, title, img string) error
	UpdateBoardPublish(userId string, boardId int, publish bool) error
}
