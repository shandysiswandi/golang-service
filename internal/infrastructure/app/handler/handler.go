package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/shandysiswandi/echo-service/internal/port"
)

/* Example get user jwt
user := c.Get("user").(*jwte.Token)
data := user.Claims.(*JWTClaim)
*/

type (
	handler struct {
		validate *validator.Validate
		tdu      port.TodoUsecase
	}

	HandlerConfig struct {
		TodoUsecase port.TodoUsecase
		Validator   *validator.Validate
	}
)

func New(hc HandlerConfig) *handler {
	return &handler{
		validate: hc.Validator,
		tdu:      hc.TodoUsecase,
	}
}
