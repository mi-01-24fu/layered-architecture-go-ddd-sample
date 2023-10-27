package model

import "fmt"

type Content struct {
	value string
}

func NewContent(value string) (*Content, error) {
	if len([]rune(value)) > 100 {
		return nil, fmt.Errorf("content is less than 100 characters Received value length: %v", len([]rune(value)))
	}
	return &Content{value: value}, nil
}

func (content *Content) String() string {
	return content.value
}
