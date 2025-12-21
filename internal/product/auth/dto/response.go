package dto

type LoginResp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
