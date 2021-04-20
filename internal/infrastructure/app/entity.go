package app

import jwte "github.com/dgrijalva/jwt-go"

type userJWT struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	CompanyID     int    `json:"company_id"`
	UserCompanyID int    `json:"user_company_id"`
}

type JWTClaim struct {
	jwte.StandardClaims
	SessID         int         `json:"sess_id"`
	User           userJWT     `json:"user"`
	Lang           interface{} `json:"lang"`
	SessionSetting int         `json:"session_setting"`
}
