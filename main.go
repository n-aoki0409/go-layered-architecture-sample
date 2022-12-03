package main

import (
	"github.com/labstack/echo"
	"github.com/n-aoki0409/go-layered-architecture-sample/config"
	"github.com/n-aoki0409/go-layered-architecture-sample/infrastructure"
	"github.com/n-aoki0409/go-layered-architecture-sample/presentation/handler"
	"github.com/n-aoki0409/go-layered-architecture-sample/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	taskRepository := infrastructure.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()
	handler.InitRouting(e, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
