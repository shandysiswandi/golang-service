package gen_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestNanoid(t *testing.T) {
	act := gen.New()
	nnid := act.Nanoid()
	assert.Equal(t, 11, len(nnid))
}

func TestUUID(t *testing.T) {
	act := gen.New()
	uuid := act.UUID()
	assert.Equal(t, 36, len(uuid))
}
