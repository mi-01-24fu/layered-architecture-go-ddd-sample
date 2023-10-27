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

	log.Println("check value empty string")
	if value == "" {
		return nil, fmt.Errorf("please enter a title")
	}

	if len([]rune(value)) < 3 {
		return nil, fmt.Errorf("title is more than 3 characters Received value length: %v", len(value))
	}

	if len([]rune(value)) > 20 {
		return nil, fmt.Errorf("title is less than 20 characters Received value length: %v", len(value))
	}

	log.Println("end NewTitle domain")
	return &Title{value: value}, nil
}

func (title *Title) String() string {
	return string(title.value)
}
