package handler

/* Example get user jwt
user := c.Get("user").(*jwte.Token)
data := user.Claims.(*JWTClaim)
*/

type (
	handler       struct{}
	HandlerConfig struct{}
)

func New(hc HandlerConfig) *handler {
	return &handler{}
}
