package model

import (
	"fmt"
	"time"
)

type Date struct {
	value string
}

const layout = "2006/01/02"

func NewDate(value string) (*Date, error) {
	if value == "" {
		return &Date{value: value}, nil
	}
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, fmt.Errorf("please format the data in yyyy/mm/dd format Received value: %v", value)
	}

	return &Date{value: t.Format("2006/01/02")}, nil
}

func (date *Date) Date() time.Time {
	t, _ := time.Parse(layout, date.value)
	return t
}

func (date *Date) String() string {
	t, _ := time.Parse(layout, date.value)
	return t.String()
}
