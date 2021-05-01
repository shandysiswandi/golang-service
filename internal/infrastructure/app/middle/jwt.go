package middle

import (
	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shandysiswandi/echo-service/internal/util/is"
)

type (
	UserJWT struct {
		ID            int    `json:"id"`
		Email         string `json:"email"`
		CompanyID     int    `json:"company_id"`
		UserCompanyID int    `json:"user_company_id"`
	}

	JWTClaim struct {
		jwtlib.StandardClaims
		SessID         int         `json:"sess_id"`
		User           UserJWT     `json:"user"`
		Lang           interface{} `json:"lang"`
		SessionSetting int         `json:"session_setting"`
	}
)

func JWT(key string, wl ...string) echo.MiddlewareFunc {
	jwtCfg := middleware.JWTConfig{
		Claims:     &JWTClaim{},
		SigningKey: []byte(key),
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return is.InArrayString(wl, path)
		},
	}
	return middleware.JWTWithConfig(jwtCfg)
}
