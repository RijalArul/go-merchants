package response

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
