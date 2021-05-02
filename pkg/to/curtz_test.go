package to_test

import (
	"testing"
	"time"

	"github.com/shandysiswandi/echo-service/pkg/to"
	"github.com/stretchr/testify/assert"
)

func TestCurrentTimezone(t *testing.T) {
	utc := time.Now().UTC()
	act := to.CurrentTimezone("Asia/Jakarta", utc)
	assert.NotEqual(t, utc, act)
	act = to.CurrentTimezone("Asi/Jakarta", utc)
	assert.Equal(t, utc, act)
}
