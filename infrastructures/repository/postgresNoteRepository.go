package repository

import (
	"assignment/commons/bootstrap"
	"assignment/domains"
	"assignment/infrastructures/sql/database"
	"context"
	"github.com/rs/xid"
)

type postgresNoteRepository struct {
	database bootstrap.Database
}

func (pnr *postgresNoteRepository) Delete(c context.Context, id string) error {
	return pnr.database.Query.DeleteNote(c, id)
}

func (pnr *postgresNoteRepository) EditNote(c context.Context, note domains.MutateNoteRequest, id string) error {
	return pnr.database.Query.EditNote(c, database.EditNoteParams{
		ID:          id,
		Title:       note.Title,
		Description: note.Body,
	})
}

func (pnr *postgresNoteRepository) Get(c context.Context, owner string) ([]domains.Note, error) {
	parsedNotes := make([]domains.Note, 0)
	notes, err := pnr.database.Query.GetNotes(c, owner)
	if err != nil {
		return parsedNotes, err
	}
	for _, note := range notes {
		parsedNotes = append(parsedNotes, domains.Note{
			ID:        note.ID,
			Title:     note.Title,
			Body:      note.Description,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
			Owner:     note.Owner,
		})
	}
	return parsedNotes, nil
}

func (pnr *postgresNoteRepository) GetById(c context.Context, id string) (domains.Note, error) {
	note, err := pnr.database.Query.GetNote(c, id)
	if err != nil {
		return domains.Note{}, err
	}

	return note.ToDomainsNote(), nil
}

func (pnr *postgresNoteRepository) Add(c context.Context, note domains.MutateNoteRequest, owner string) (domains.MutateNoteResponseData, error) {
	id := xid.New().String()
	err := pnr.database.Query.CreateNote(c, database.CreateNoteParams{
		ID:          id,
		Title:       note.Title,
		Description: note.Body,
		Owner:       owner,
	})
	if err != nil {
		return domains.MutateNoteResponseData{}, err
	}
	response := domains.MutateNoteResponseData{
		ID:    id,
		Title: note.Title,
		Owner: owner,
	}
	return response, nil
}

func NewPostgresNoteRepository(database bootstrap.Database) domains.NoteRepository {
	return &postgresNoteRepository{
		database: database,
	}
}
