package model

import (
	"fmt"
	"log"
)

type Title struct {
	value string
}

func NewTitle(value string) (*Title, error) {
	log.Println("start NewTitle domain")

	strlen := len([]rune(value))

	log.Println("check value empty string")
	if value == "" {
		return nil, fmt.Errorf("please enter a title")
	}

	if strlen < 3 {
		return nil, fmt.Errorf("title should be at least 3 characters. Received value length: %v", strlen)
	}

	if strlen > 20 {
		return nil, fmt.Errorf("title should be less than 20 characters. Received value length: %v", strlen)
	}

	log.Println("end NewTitle domain")
	return &Title{value: value}, nil
}

func (title *Title) String() string {
	return title.value
}
