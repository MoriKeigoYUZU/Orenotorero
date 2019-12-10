package requestBody

import "time"

type CardChangeTitle struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type CardChangeDeadline struct {
	Id       int       `json:"id"`
	Deadline time.Time `json:"deadline"`
}

type CardChangePosition struct {
	Id          int `json:"id"`
	Current     int `json:"current_position"`
	Destination int `json:"destination_position"`
}

type CardCreate struct {
	Title    string `json:"title"`
	KanbanId int    `json:"kanban_id"`
	Position int    `json:"position"`
}

type CardAddFile struct {
	Id       int    `json:"id"`
	FileName string `json:"name"`
	FileData string `json:"data"`
}