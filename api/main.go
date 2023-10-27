package main

import (
	"github.com/labstack/echo/v4"
	"layered-architecture-go-ddd-sample/config"
	"layered-architecture-go-ddd-sample/domain/model"
	"layered-architecture-go-ddd-sample/infrastructure"
	"layered-architecture-go-ddd-sample/presentation/handler"
	"layered-architecture-go-ddd-sample/usecase"
)

func main() {
	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	memoFactory := model.NewMemoFactoryImpl()
	memoRepository := infrastructure.NewMemoRepository(db, memoFactory)
	memoUsecase := usecase.NewMemoUsecase(memoFactory, memoRepository)
	memoHandler := handler.NewMemoHandler(memoUsecase)

	e := echo.New()
	handler.InitRouting(e, memoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
