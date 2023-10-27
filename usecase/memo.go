package usecase

import (
	"context"
	"fmt"
	"layered-architecture-go-ddd-sample/domain/model"
	"layered-architecture-go-ddd-sample/domain/model/repository"
	"log"
)

type MemoUsecase interface {
	Create(ctx context.Context, title, content, date string) (*model.Memo, error)
	Get(ctx context.Context) ([]*model.Memo, error)
}

type memoUsecase struct {
	memoFactory    model.MemoFactory
	memoRepository repository.MemoRepository
}

func NewMemoUsecase(memoFactory model.MemoFactory, memoRepository repository.MemoRepository) MemoUsecase {
	return &memoUsecase{
		memoFactory:    memoFactory,
		memoRepository: memoRepository,
	}
}

func (m memoUsecase) Create(ctx context.Context, title, content, date string) (*model.Memo, error) {
	log.Println("start create usecase")

	log.Println("call memo factory create method")
	memo, err := m.memoFactory.Create(title, content, date)
	if err != nil {
		return nil, err
	}

	log.Println("call memo repository count method")
	isMemoExists, err := m.memoRepository.Exists(ctx, memo)
	if err != nil {
		return nil, err
	}

	if isMemoExists {
		return nil, fmt.Errorf("same memo exists Received value: %v", title)
	}

	log.Println("call memo repository create method")
	createdMemo, err := m.memoRepository.Create(ctx, memo)
	if err != nil {
		return nil, err
	}

	log.Println("end create usecase")
	return createdMemo, nil
}

func (m memoUsecase) Get(ctx context.Context) ([]*model.Memo, error) {
	memos, err := m.memoRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	return memos, nil
}
