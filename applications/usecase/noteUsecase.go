package usecase

import (
	"assignment/domains"
	"context"
	"time"
)

type noteUseCase struct {
	noteRepository domains.NoteRepository
	contextTimeout time.Duration
}

func (nu *noteUseCase) DeleteNote(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, nu.contextTimeout)
	defer cancel()
	return nu.noteRepository.Delete(ctx, id)
}

func (nu *noteUseCase) EditNote(c context.Context, note domains.MutateNoteRequest, id string) error {
	ctx, cancel := context.WithTimeout(c, nu.contextTimeout)
	defer cancel()
	return nu.noteRepository.EditNote(ctx, note, id)
}

func (nu *noteUseCase) Get(c context.Context, owner string) ([]domains.Note, error) {
	ctx, cancel := context.WithTimeout(c, nu.contextTimeout)
	defer cancel()
	return nu.noteRepository.Get(ctx, owner)
}

func (nu *noteUseCase) GetById(c context.Context, id string) (domains.Note, error) {
	ctx, cancel := context.WithTimeout(c, nu.contextTimeout)
	defer cancel()
	return nu.noteRepository.GetById(ctx, id)
}

func (nu *noteUseCase) Add(c context.Context, task domains.MutateNoteRequest, owner string) (domains.MutateNoteResponseData, error) {
	ctx, cancel := context.WithTimeout(c, nu.contextTimeout)
	defer cancel()
	return nu.noteRepository.Add(ctx, task, owner)
}

func NewNoteUsecase(taskRepository domains.NoteRepository, timeout time.Duration) domains.NoteUsecase {
	return &noteUseCase{
		noteRepository: taskRepository,
		contextTimeout: timeout,
	}
}
