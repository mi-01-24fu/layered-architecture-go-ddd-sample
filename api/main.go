package main

import (
	"github.com/labstack/echo/v4"
	"layered-architecture-go-ddd-sample/config"
	"layered-architecture-go-ddd-sample/domain/model"
	"layered-architecture-go-ddd-sample/infrastructure"
	"layered-architecture-go-ddd-sample/interface/handler"
	"layered-architecture-go-ddd-sample/usecase"
)

func main() {
	db := config.NewDB()
	defer db.Close()
	memoRepository := infrastructure.NewMemoRepository(db)
	memoFactory := model.NewMemoFactoryImpl()
	memoUsecase := usecase.NewMemoUsecase(memoFactory, memoRepository)
	memoHandler := handler.NewMemoHandler(memoUsecase)

	e := echo.New()
	handler.InitRouting(e, memoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
