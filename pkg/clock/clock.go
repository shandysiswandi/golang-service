package clock

import "time"

type (
	Clocker interface {
		Now() time.Time
	}

	clock struct{}
)

func New() Clocker {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now()
}
