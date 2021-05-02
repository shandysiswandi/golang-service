package times

import "time"

type (
	Times interface {
		Now() time.Time
	}
	times struct{}
)

func New() Times {
	return &times{}
}

func (*times) Now() time.Time {
	return time.Now()
}
