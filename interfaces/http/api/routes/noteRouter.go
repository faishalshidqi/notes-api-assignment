package routes

import (
	"assignment/applications/usecase"
	"assignment/commons/bootstrap"
	"assignment/infrastructures/repository"
	"assignment/infrastructures/security"
	"assignment/interfaces/http/api/controllers"
	"github.com/gin-gonic/gin"
	"time"
)

func newNoteRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	noteRepository := repository.NewMysqlNoteRepository(db)
	tokenManager := security.NewJwtTokenManager()
	noteController := controllers.NoteController{
		NoteUsecase:  usecase.NewNoteUsecase(noteRepository, timeout),
		TokenManager: tokenManager,
		Env:          env,
	}
	router.POST("/notes", noteController.AddNote)
	router.GET("/notes/", noteController.GetNotes)
	router.GET("/notes/:note_id", noteController.GetNote)
	router.PUT("/notes/:note_id", noteController.EditNote)
	router.DELETE("/notes/:note_id", noteController.DeleteNote)
}
