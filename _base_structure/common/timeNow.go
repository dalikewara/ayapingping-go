package common

import "time"

func TimeNowUTC() time.Time {
	return time.Now().UTC()
}
