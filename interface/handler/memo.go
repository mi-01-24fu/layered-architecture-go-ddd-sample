package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"layered-architecture-go-ddd-sample/usecase"
	"log"
	"net/http"
)

type MemoHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type memoHandler struct {
	memousecase usecase.MemoUsecase
}

func NewMemoHandler(memousecase usecase.MemoUsecase) MemoHandler {
	return &memoHandler{memousecase: memousecase}
}

type requestMemo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type responseMemo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

func (m *memoHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("------------------------------")
		log.Println("start post handler")

		log.Println("request bind")
		var req requestMemo
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		log.Println("call memo usecase create method")
		createdMemo, err := m.memousecase.Create(context.Background(), req.Title, req.Title, req.Date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMemo{
			ID:      createdMemo.Id,
			Title:   createdMemo.Title.String(),
			Content: createdMemo.Content.String(),
			Date:    createdMemo.Date.String(),
		}

		log.Println("end post handler")
		log.Println("------------------------------")
		return c.JSON(http.StatusCreated, res)
	}
}

func (m memoHandler) Get() echo.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (m memoHandler) Put() echo.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (m memoHandler) Delete() echo.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
