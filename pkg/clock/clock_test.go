package clock_test

import (
	"testing"
	"time"

	"github.com/shandysiswandi/echo-service/pkg/clock"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	n := clock.New().Now()
	assert.NotEqual(t, time.Now(), n)
}
