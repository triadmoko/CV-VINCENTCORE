package helpers

import (
	"errors"
	"time"

	readingtime "github.com/begmaroman/reading-time"
)

func FormatTimeToDate(time time.Time) string {
	format := time.Format("2006-01-02")
	return time.Format(format)
}
func ReadingTime(str string) string {
	estimation := readingtime.Estimate(str)
	return estimation.Text
}

func RangeHour(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	end = end.AddDate(0, 0, 1)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.Add(time.Hour * 1)
		return date
	}
}

func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func RangeWeek(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 7)
		return date
	}
}

func RangeMonth(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 1, 0)
		return date
	}
}

func StringToTimestamp(input string) (*time.Time, error) {
	layout := "2006-01-02"
	date, err := time.Parse(layout, input)
	if err != nil {
		return nil, errors.New("failed format date")
	}
	return &date, nil
}

func TimestampToString(input time.Time) string {
	layout := "2006-01-02"
	date := input.Format(layout)
	return date
}

func TimestampAdd(input time.Time, add time.Duration) string {
	timeAdd := input.Add(add * (24 * time.Hour))
	time := TimestampToString(timeAdd)
	return time
}
