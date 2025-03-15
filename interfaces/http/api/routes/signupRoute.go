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

func newSignupRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	userRepository := repository.NewPostgresUserRepository(db)
	passwordHash := security.NewBcryptPasswordHash()
	signupController := controllers.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepository, timeout),
		PasswordHash:  passwordHash,
		Env:           env,
	}
	router.POST("/users", signupController.Signup)
	router.GET("/users/:user_id", signupController.GetUser)
}
