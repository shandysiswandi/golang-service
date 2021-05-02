package domain_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	model := domain.Todo{}
	tb := model.TableName()

	assert.Equal(t, "todos", tb)
}
