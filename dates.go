package main

import (
	"fmt"
	"time"
)

type DateObject struct {
	Date      string // YYYY-MM-DD
	Title     string // Mon (12/10)
	Yesterday string // YYYY-MM-DD
	Tomorrow  string // YYYY-MM-DD
}

func getDateForOffset(offset int) DateObject {
	o := offset + *initialOffset
	date := time.Now().Add(time.Duration(o) * 24 * time.Hour)

	return DateObject{
		Date:      date.Format(time.DateOnly),
		Title:     fmt.Sprintf("%s (%d/%d)", date.Weekday().String(), date.Day(), date.Month()),
		Yesterday: date.Add(-1 * 24 * time.Hour).Format(time.DateOnly),
		Tomorrow:  date.Add(24 * time.Hour).Format(time.DateOnly),
	}
}
