package request

type CustomerLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
