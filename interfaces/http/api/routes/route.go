package routes

import (
	"assignment/commons/bootstrap"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, gin *gin.Engine) {
	router := gin.Group("")
	newSignupRouter(env, timeout, db, router)
	newAuthnRouter(env, timeout, db, router)

	newNoteRouter(env, timeout, db, router)
}
