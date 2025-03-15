package domains

import (
	"context"
	"time"
)

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Body      string    `json:"body" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Owner     string    `json:"owner"`
}

type MutateNoteRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type MutateNoteResponse struct {
	Message string                 `json:"message"`
	Status  string                 `json:"status"`
	Data    MutateNoteResponseData `json:"data"`
}

type MutateNoteResponseData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
}

type GetNoteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Note   `json:"data"`
}

type GetNotesResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []Note `json:"data"`
}

type NoteRepository interface {
	Add(c context.Context, note MutateNoteRequest, owner string) (MutateNoteResponseData, error)
	Get(c context.Context, owner string) ([]Note, error)
	GetById(c context.Context, id string) (Note, error)
	EditNote(c context.Context, note MutateNoteRequest, id string) error
	Delete(c context.Context, id string) error
}

type NoteUsecase interface {
	Add(c context.Context, task MutateNoteRequest, owner string) (MutateNoteResponseData, error)
	Get(c context.Context, owner string) ([]Note, error)
	GetById(c context.Context, id string) (Note, error)
	EditNote(c context.Context, note MutateNoteRequest, id string) error
	DeleteNote(c context.Context, id string) error
}
