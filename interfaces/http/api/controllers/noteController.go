package controllers

import (
	"assignment/applications/security"
	"assignment/commons/bootstrap"
	"assignment/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NoteController struct {
	NoteUsecase  domains.NoteUsecase
	TokenManager security.AuthnTokenManager
	Env          *bootstrap.Env
}

// AddNote Create A New Thread godoc
//
//	@Summary	Create Note
//	@Body		Creating a new note. Only valid users can create a note
//	@Tags		notes
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer Token"
//	@Param		title			body		string	true	"title of the note"
//	@Param		description		body		string	true	"body of the note"
//	@Success	201				{object}	domains.MutateNoteResponse
//	@Failure	400				{object}	domains.ErrorResponse
//	@Failure	401				{object}	domains.ErrorResponse
//	@Failure	500				{object}	domains.ErrorResponse
//	@Router		/notes [post]
func (nc *NoteController) AddNote(c *gin.Context) {
	token, err := nc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid token structure",
		})
		return
	}
	id, err := nc.TokenManager.VerifyToken(token, nc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	addThreadRequest := domains.MutateNoteRequest{}
	if err := c.ShouldBind(&addThreadRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Failed to bind request body",
		})
		return
	}
	addedNote, err := nc.NoteUsecase.Add(c, addThreadRequest, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: "Failed to add note",
		})
		return
	}
	c.JSON(http.StatusCreated, domains.MutateNoteResponse{
		Message: "successfully added note",
		Status:  "success",
		Data:    addedNote,
	})
}

// DeleteNote Delete A Note By ID godoc
//
//	@Summary	Delete Note
//	@Body		Deleting a note. Only note's owner can delete a note
//	@Tags		notes
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer Token"
//	@Param		note_id			path		string	true	"id of the note"
//	@Success	200				{object}	domains.SuccessResponse
//	@Failure	400				{object}	domains.ErrorResponse
//	@Failure	401				{object}	domains.ErrorResponse
//	@Failure	403				{object}	domains.ErrorResponse
//	@Failure	404				{object}	domains.ErrorResponse
//	@Failure	500				{object}	domains.ErrorResponse
//	@Router		/notes/{note_id} [delete]
func (nc *NoteController) DeleteNote(c *gin.Context) {
	token, err := nc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	id, err := nc.TokenManager.VerifyToken(token, nc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	noteId := c.Param("note_id")
	note, err := nc.NoteUsecase.GetById(c, noteId)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Status:  "fail",
			Message: "Note not found",
		})
		return
	}
	if id != note.Owner {
		c.JSON(http.StatusForbidden, domains.ErrorResponse{
			Status:  "fail",
			Message: "Unauthorized access",
		})
		return
	}
	err = nc.NoteUsecase.DeleteNote(c, noteId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: "Failed to edit note",
		})
		return
	}
	c.JSON(http.StatusOK, domains.SuccessResponse{
		Message: "successfully deleted note",
		Status:  "success",
	})
}

// EditNote Edit A Note By ID godoc
//
//	@Summary	Edit Note
//	@Body		Editing a note. Only note's owner can delete a note
//	@Tags		notes
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer Token"
//	@Param		note_id			path		string	true	"id of the note"
//	@Success	200				{object}	domains.SuccessResponse
//	@Failure	400				{object}	domains.ErrorResponse
//	@Failure	401				{object}	domains.ErrorResponse
//	@Failure	403				{object}	domains.ErrorResponse
//	@Failure	404				{object}	domains.ErrorResponse
//	@Failure	500				{object}	domains.ErrorResponse
//	@Router		/notes/{note_id} [put]
func (nc *NoteController) EditNote(c *gin.Context) {
	token, err := nc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	id, err := nc.TokenManager.VerifyToken(token, nc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	addThreadRequest := domains.MutateNoteRequest{}
	if err := c.ShouldBind(&addThreadRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Failed to bind request body",
		})
		return
	}
	noteId := c.Param("note_id")
	note, err := nc.NoteUsecase.GetById(c, noteId)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Status:  "fail",
			Message: "Note not found",
		})
		return
	}
	if id != note.Owner {
		c.JSON(http.StatusForbidden, domains.ErrorResponse{
			Status:  "fail",
			Message: "Unauthorized access",
		})
		return
	}
	err = nc.NoteUsecase.EditNote(c, addThreadRequest, noteId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: "Failed to edit note",
		})
		return
	}
	c.JSON(http.StatusOK, domains.SuccessResponse{
		Message: "successfully edited note",
		Status:  "success",
	})
}

// GetNotes Get Notes By User ID godoc
//
//	@Summary	Get All Notes Owned By The User Logged In
//	@Body		Fetching notes. Only notes owner can see their own notes
//	@Tags		notes
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer Token"
//	@Success	200				{object}	domains.GetNotesResponse
//	@Failure	401				{object}	domains.ErrorResponse
//	@Failure	404				{object}	domains.ErrorResponse
//	@Failure	500				{object}	domains.ErrorResponse
//	@Router		/notes [get]
func (nc *NoteController) GetNotes(c *gin.Context) {
	token, err := nc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	id, err := nc.TokenManager.VerifyToken(token, nc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	notes, err := nc.NoteUsecase.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Status:  "fail",
			Message: "Failed to get notes",
		})
		return
	}
	c.JSON(http.StatusOK, domains.GetNotesResponse{
		Status:  "success",
		Message: "Successfully get notes",
		Data:    notes,
	})
}

// GetNote Get Note By Note ID godoc
//
//	@Summary	Get Note By ID
//	@Body		Fetching note by id. Only notes owner can see their own notes
//	@Tags		notes
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer Token"
//	@Param		note_id			path		string	true	"id of the note"
//	@Success	200				{object}	domains.GetNoteResponse
//	@Failure	401				{object}	domains.ErrorResponse
//	@Failure	403				{object}	domains.ErrorResponse
//	@Failure	404				{object}	domains.ErrorResponse
//	@Failure	500				{object}	domains.ErrorResponse
//	@Router		/notes/{note_id} [get]
func (nc *NoteController) GetNote(c *gin.Context) {
	token, err := nc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	id, err := nc.TokenManager.VerifyToken(token, nc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid bearer token",
		})
		return
	}
	noteId := c.Param("note_id")
	note, err := nc.NoteUsecase.GetById(c, noteId)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Status:  "fail",
			Message: "Note does not exist",
		})
		return
	}
	if id != note.Owner {
		c.JSON(http.StatusForbidden, domains.ErrorResponse{
			Status:  "fail",
			Message: "Unauthorized access",
		})
		return
	}
	c.JSON(http.StatusOK, domains.GetNoteResponse{
		Status:  "success",
		Message: "Note fetched successfully",
		Data:    note,
	})
}
