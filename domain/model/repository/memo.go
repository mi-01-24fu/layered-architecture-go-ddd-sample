package repository

import (
	"context"
	"layered-architecture-go-ddd-sample/domain/model"
)

type MemoRepository interface {
	Count(ctx context.Context, memo *model.Memo) (bool, error)
	Create(ctx context.Context, memo *model.Memo) (*model.Memo, error)
}
