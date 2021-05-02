package times_test

import (
	"testing"
	"time"

	"github.com/shandysiswandi/echo-service/pkg/times"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	tm := times.New().Now()
	assert.NotEqual(t, time.Now(), tm)
}
