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

func newAuthnRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	userRepository := repository.NewPostgresUserRepository(db)
	tokenManager := security.NewJwtTokenManager()
	passwordHash := security.NewBcryptPasswordHash()
	authenticationController := controllers.AuthenticationController{
		AuthenticationUsecase: usecase.NewAuthenticationUsecase(userRepository, tokenManager, passwordHash, timeout),
		Env:                   env,
	}
	router.POST("/authentications", authenticationController.Login)
}
