package is_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/is"
	"github.com/stretchr/testify/assert"
)

func Test_InArrayString(t *testing.T) {
	ts := []struct {
		name     string
		inputOne []string
		inputTwo []string
		expected bool
	}{
		{"one", []string{}, []string{}, false},
		{"two", []string{"satu", "dua"}, []string{}, false},
		{"two", []string{"satu", "dua"}, []string{"tiga"}, false},
		{"two", []string{"satu", "dua"}, []string{"satu"}, true},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			act := is.InArrayString(tt.inputOne, tt.inputTwo...)
			assert.Equal(t, tt.expected, act)
		})
	}
}
