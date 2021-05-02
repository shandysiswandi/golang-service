package validation_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/validation"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	act := validation.New()
	assert.NotNil(t, act)
}

func TestValidate(t *testing.T) {
	// setup
	type Testing struct {
		Field string `json:"field" validate:"required"`
	}
	var test Testing
	act := validation.New()

	// errro 1
	err := act.Validate(nil)
	assert.Error(t, err)

	// error 2
	err = act.Validate(test)
	assert.Error(t, err)

	// pass validation
	test.Field = "kuda"
	err = act.Validate(test)
	assert.NoError(t, err)
}

func TestValidateVar(t *testing.T) {
	// setup
	act := validation.New()

	// errro 1
	err := act.ValidateVar(nil, "required,min=5,email")
	assert.Error(t, err)

	// error required
	err = act.ValidateVar("", "required")
	assert.Error(t, err)

	// error min
	err = act.ValidateVar("aaaa", "min=5")
	assert.Error(t, err)

	// error email
	err = act.ValidateVar("a-a.com", "email")
	assert.Error(t, err)

	// pass validation
	err = act.ValidateVar("a@a.com", "required,min=5,email")
	assert.NoError(t, err)
}

func TestExtractErrorMessage(t *testing.T) {
	// setup
	type Testing struct {
		Email string `json:"email" validate:"email"`
		Name  string `json:"name" validate:"min=5"`
		Title string `json:"title" validate:"required"`
		UUID  string `json:"uuid" validate:"uuid"`
	}
	test := Testing{}
	act := validation.New()
	err := act.Validate(test)
	errVar := act.ValidateVar("ku", "email")
	errVarInvalid := act.ValidateVar(nil, "email")

	// testing
	me := act.ExtractErrorMessage(err)
	meVar := act.ExtractErrorMessage(errVar)
	meVarInvalid := act.ExtractErrorMessage(errVarInvalid)

	// assertion
	assert.NotEmpty(t, me)
	assert.Equal(t, 4, len(me))

	// assertion
	assert.NotEmpty(t, meVar)
	assert.Equal(t, 1, len(meVar))

	// assertion
	assert.NotEmpty(t, meVarInvalid)
	assert.Equal(t, 1, len(meVarInvalid))
}
