package dto

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

type Response struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
