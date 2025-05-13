package dto

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

type AuthDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
}

type UserSignupDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
