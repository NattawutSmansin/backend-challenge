package initialize

import (
	"backend-challenge/module/thirdChallenge/handler"
	"backend-challenge/module/thirdChallenge/repository"
	"backend-challenge/module/thirdChallenge/usecase"
)

func InitializeHandler() *handler.Handler {
	repo := repository.NewRepository()
	usecase := usecase.NewUseCase(repo)
	return handler.NewHandler(usecase)
}
