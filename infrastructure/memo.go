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
)

type MemoRepository struct {
	DB          *sql.DB
	MemoFactory model.MemoFactory
}

func NewMemoRepository(db *sql.DB, memoFactory model.MemoFactory) repository.MemoRepository {
	return &MemoRepository{
		DB:          db,
		MemoFactory: memoFactory,
	}
}

func (m MemoRepository) Exists(ctx context.Context, memo *model.Memo) (bool, error) {
	count, err := models.Memos(models.MemoWhere.Title.EQ(memo.Title.String())).Count(ctx, m.DB)
	if err != nil {
		return true, fmt.Errorf("duplicate confirmation error: %v", err)
	}
	if count != 0 {
		return true, nil
	}

	return false, nil
}

func (m MemoRepository) Create(ctx context.Context, memo *model.Memo) (*model.Memo, error) {
	input := &models.Memo{
		Title:   memo.Title.String(),
		Content: null.StringFrom(memo.Content.String()),
		Date: null.Time{
			Time: memo.Date.Date(),
		},
	}

	err := input.Insert(ctx, m.DB, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("database insert error: %v", err)
	}

	return memo, nil
}

func (m MemoRepository) Get(ctx context.Context) ([]*model.Memo, error) {
	output, err := models.Memos().All(ctx, m.DB)
	if err != nil {
		return nil, fmt.Errorf("database get all error: %v", err)
	}

	var memos []*model.Memo

	for _, memo := range output {
		createdMemo, err := m.MemoFactory.Create(memo.Title, memo.Content.String, memo.Date.Time.String())
		if err != nil {
			return nil, err
		}
		createdMemo.SetID(memo.ID)
		memos = append(memos, createdMemo)
	}

	return memos, nil
}
