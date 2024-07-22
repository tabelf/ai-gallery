package utils

import (
	"time"
)

func ParseYMD(timestamp string) (time.Time, error) {
	t, err := time.Parse(time.DateOnly, timestamp)
	if err != nil {
		return t, err
	}
	return t, nil
}

func FirstTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func LastTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}
