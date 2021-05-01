package middle

import (
	"time"

	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shandysiswandi/echo-service/pkg/is"
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

func CreateJWTToken(key string) (string, error) {
	// Set custom claims
	claims := &JWTClaim{
		SessID: time.Now().Nanosecond(),
		User: UserJWT{
			ID:            1,
			Email:         "",
			CompanyID:     1,
			UserCompanyID: 1,
		},
		Lang:           nil,
		SessionSetting: 1,
		StandardClaims: jwtlib.StandardClaims{
			Audience:  "audien",
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			Id:        "id",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "issuer",
			NotBefore: 0,
			Subject:   "subject",
		},
	}

	// Create token with claims
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}
