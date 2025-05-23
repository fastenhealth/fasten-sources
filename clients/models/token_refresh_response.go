package models

type TokenRefreshResponse struct {
	//only used for confidential clients
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	IdToken      string `json:"id_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	PatientId    string `json:"patient,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
}
