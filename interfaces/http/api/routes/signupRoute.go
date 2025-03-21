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
	userRepository := repository.NewMysqlUserRepository(db)
	passwordHash := security.NewBcryptPasswordHash()
	tokenManager := security.NewJwtTokenManager()
	signupController := controllers.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepository, timeout),
		PasswordHash:  passwordHash,
		TokenManager:  tokenManager,
		Env:           env,
	}
	router.POST("/users", signupController.Signup)
	router.GET("/users", signupController.GetUser)
}
