package model

import "fmt"

type Content struct {
	value string
}

func NewContent(value string) (*Content, error) {
	if len(value) > 100 {
		return nil, fmt.Errorf("title is less than 100 characters Received value length: %v", len(value))
	}
	return &Content{value: value}, nil
}

func (content *Content) String() string {
	return content.value
}
