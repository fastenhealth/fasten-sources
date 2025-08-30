package models

type TokenIntrospectResponse struct {
	Active    bool   `json:"active,omitempty"`
	Scope     string `json:"scope,omitempty"`
	ClientID  string `json:"client_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Patient   string `json:"patient,omitempty"`
	ExpiresAt int    `json:"exp,omitempty"`
}
