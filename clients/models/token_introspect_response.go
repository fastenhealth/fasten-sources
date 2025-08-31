package models

type TokenIntrospectResponse struct {
	Active    bool   `json:"active,omitempty"`
	Scope     string `json:"scope,omitempty"`
	ClientID  string `json:"client_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Patient   string `json:"patient,omitempty"`
	ExpiresAt int    `json:"exp,omitempty"`
}

// https://datatracker.ietf.org/doc/html/rfc7009#section-4.1.2.2
type TokenIntrospectTokenType string

const (
	TokenIntrospectTokenTypeAccess  TokenIntrospectTokenType = "access_token"
	TokenIntrospectTokenTypeRefresh TokenIntrospectTokenType = "refresh_token"
)
