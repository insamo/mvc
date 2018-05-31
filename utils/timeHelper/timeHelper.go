package timeHelper

import (
	"time"
)

func FewDaysLater(day int) time.Time {
	return FewDurationLater(time.Duration(day) * 24 * time.Hour)
}

func TwentyFourHoursLater() time.Time {
	return FewDurationLater(time.Duration(24) * time.Hour)
}

func SixHoursLater() time.Time {
	return FewDurationLater(time.Duration(6) * time.Hour)
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func InTimeSpanNow(start, end time.Time) bool {
	now := time.Now()
	return InTimeSpan(start, end, now)
}

func FewDurationLater(duration time.Duration) time.Time {
	// When Save time should considering UTC
	// baseTime := time.Now()
	// log.Debugf("basetime : %s", baseTime)
	fewDurationLater := time.Now().Add(duration)
	return fewDurationLater
}

func FewDurationLaterMillisecond(duration time.Duration) int64 {
	return FewDurationLater(duration).UnixNano() / int64(time.Millisecond)
}

func IsExpired(expirationTime time.Time) bool {
	// baseTime := time.Now()
	// log.Debugf("basetime : %s", baseTime)
	// elapsed := time.Since(expirationTime)
	// log.Debugf("elapsed : %s", elapsed)
	after := time.Now().After(expirationTime)
	return after
}

func SetTime(datetime time.Time, h, m, s int) time.Time {
	t := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), h, m, s, 0, datetime.Location())
	return t
}
