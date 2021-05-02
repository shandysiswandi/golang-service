package to

import "time"

func CurrentTimezone(tz string, t time.Time) time.Time {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return t
	}
	return t.In(loc)
}
