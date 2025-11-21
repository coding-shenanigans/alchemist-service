package dto

import "net/http"

type UserSession struct {
	Email         string       `json:"email"`
	Username      string       `json:"username"`
	AccessToken   string       `json:"accessToken"`
	SessionCookie *http.Cookie `json:"-"`
}
