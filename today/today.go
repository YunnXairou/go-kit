// Package today provides utilities for access regarding today's date.
package today

import "time"

var (
	maxNano = 999999999
)

// TodayBeforeMidnight returns the current date with time set to directly before midnight.
// For example, 2016-06-24 11:59:59.999
func TodayBeforeMidnight() time.Time {
	year, month, day := time.Now().Date()

	return time.Date(year, month, day, 23, 59, 59, maxNano, time.Local)
}