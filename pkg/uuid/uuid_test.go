package uuid_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_New_And_Generate(t *testing.T) {
	gen := uuid.Generate()

	assert.NotEqual(t, "", gen)
	assert.Equal(t, 36, len(gen))
}
