package model

import "time"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Session struct {
	Token     string
	ExpiresAt time.Time
}
