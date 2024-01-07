package timeLib

import "time"

// NowUTC returns time now in UTC
func NowUTC() time.Time {
	return time.Now().UTC()
}
