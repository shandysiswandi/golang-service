package nanoid_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/nanoid"
	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	id, err := nanoid.Generate()
	assert.NoError(t, err)

	assert.NotEqual(t, "", id)
	assert.Equal(t, 21, len(id))
}
