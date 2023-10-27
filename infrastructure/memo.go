package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"layered-architecture-go-ddd-sample/domain/model"
	"layered-architecture-go-ddd-sample/domain/model/repository"
	"layered-architecture-go-ddd-sample/infrastructure/models"
	"log"
)

type MemoRepository struct {
	DB *sql.DB
}

func NewMemoRepository(DB *sql.DB) repository.MemoRepository {
	return &MemoRepository{DB: DB}
}

func (m MemoRepository) Count(ctx context.Context, memo *model.Memo) (bool, error) {
	log.Println("start count infrastructure")

	log.Println("call models count method")
	count, err := models.Memos(models.MemoWhere.Title.EQ(memo.Title.String())).Count(ctx, m.DB)
	if err != nil {
		return true, fmt.Errorf("duplicate confirmation error: %v", err)
	}
	if count != 0 {
		return true, nil
	}

	log.Println("end count infrastructure")
	return false, nil
}

func (m MemoRepository) Create(ctx context.Context, memo *model.Memo) (*model.Memo, error) {
	log.Println("start create infrastructure")

	input := &models.Memo{
		ID:      memo.Id,
		Title:   memo.Title.String(),
		Content: null.StringFrom(memo.Content.String()),
		Date: null.Time{
			Time: memo.Date.Date(),
		},
	}

	log.Println("call models insert method")
	err := input.Insert(ctx, m.DB, boil.Infer())
	if err != nil {
		return &model.Memo{}, fmt.Errorf("database insert error: %v", err)
	}

	log.Println("end create infrastructure")
	return memo, nil
}
