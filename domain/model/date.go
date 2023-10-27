package model

import (
	"fmt"
	"regexp"
	"time"
)

type Date struct {
	value string
}

const (
	layout   = "2006/01/02"
	dbLayout = "2006-01-02 15:04:05 -0700 MST"
)

func NewDate(value string) (*Date, error) {
	if value == "" || value == "0001-01-01" {
		return &Date{}, nil
	}

	t, err := time.Parse(dbLayout, value)
	if err == nil {
		return &Date{value: t.Format("2006/01/02")}, nil
	}

	t, err = time.Parse(layout, value)
	if err != nil {
		return nil, fmt.Errorf("please format the date in yyyy/mm/dd or DB format. Received value: %v", value)
	}

	return &Date{value: t.Format("2006/01/02")}, nil
}

func check(dateStr string) bool {
	reg := regexp.MustCompile(`[-|/|:| |　]`)

	str := reg.ReplaceAllString(dateStr, "")

	format := string([]rune("20060102150405")[:len(str)])

	_, err := time.Parse(format, str)
	return err == nil
}

func (date *Date) Date() time.Time {
	t, err := time.Parse(layout, date.value)
	if err != nil {
		// TODO
	}
	return t
}

func (date *Date) String() string {
	_, err := time.Parse(layout, date.value)
	if err != nil {
		// TODO
	}
	return date.value
}
