package handler

import (
	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/port"
	"github.com/shandysiswandi/echo-service/pkg/validation"
)

/* Example get user jwt
user := c.Get("user").(*jwte.Token)
data := user.Claims.(*JWTClaim)
*/

type (
	handler struct {
		config   *config.Config
		validate *validation.Validation
		tdu      port.TodoUsecase
	}

	HandlerConfig struct {
		Config      *config.Config
		Validator   *validation.Validation
		TodoUsecase port.TodoUsecase
	}
)

func New(hc HandlerConfig) *handler {
	return &handler{
		config:   hc.Config,
		validate: hc.Validator,
		tdu:      hc.TodoUsecase,
	}
}
