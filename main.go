package main

import (
	"go-echo-api/controller"
	"go-echo-api/db"
	"go-echo-api/repository"
	"go-echo-api/router"
	"go-echo-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
