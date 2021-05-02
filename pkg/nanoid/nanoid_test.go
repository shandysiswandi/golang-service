package nanoid_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/nanoid"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	id := nanoid.New().Generate()

	assert.NotEqual(t, "", id)
	assert.Equal(t, 11, len(id))
}
