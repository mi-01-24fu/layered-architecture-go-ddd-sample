package model

import "log"

type MemoFactory interface {
	Create(title, content, date string) (*Memo, error)
}

type MemoFactoryImpl struct{}

func NewMemoFactoryImpl() *MemoFactoryImpl {
	return &MemoFactoryImpl{}
}

func (m MemoFactoryImpl) Create(title, content, date string) (*Memo, error) {
	log.Println("start create factory")

	createTitle, err := NewTitle(title)
	if err != nil {
		return nil, err
	}

	createContent, err := NewContent(content)
	if err != nil {
		return nil, err
	}

	createDate, err := NewDate(date)
	if err != nil {
		return nil, err
	}

	memo := NewMemo(*createTitle, *createContent, *createDate)
	return memo, nil
}
