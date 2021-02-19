// Package vendors ..
package vendors

import (
	"time"
)

// StartOfDay => // Return the Start of The Day
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// EndOfDay => // Return The End Of Day
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 59, time.UTC)
}

// BetwenToday => // Return StartOfToday, EndOfToday
func BetwenToday() (time.Time, time.Time) {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), time.Date(year, month, day, 23, 59, 59, 59, time.UTC)
}

// BetwenYesterDay => // Return StartOfToday, EndOfToday
func BetwenYesterDay() (time.Time, time.Time) {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	day = day - 1
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), time.Date(year, month, day, 23, 59, 59, 59, time.UTC)
}

// BetwenLastSevenDay ..
func BetwenLastSevenDay() (time.Time, time.Time) {
	now := time.Now()

	year, month, day := now.Date()

	startOfLastSevenDay := time.Date(year, month, day-7, 0, 0, 0, 0, time.UTC)
	toToday := time.Date(year, month, day, 23, 59, 59, 59, time.UTC)

	return startOfLastSevenDay, toToday
}
