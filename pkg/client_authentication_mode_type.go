package pkg

// ClientAuthenticationMethodType - Supported Client Authentication Methods
// FHIR servers may support different client authentication methods for the OAuth applications that Fasten has registered.
// This enum is used to differentiate between the various client authentication methods that we need to support in our code
// Different client authentication methods will require different methods for refreshing the access token.
// This library supports:
// - client_secret_basic: (default) uses the HTTP `Basic` authentication scheme to send `client_id` and `client_secret` in an `Authorization` HTTP Header.
// - private_key_jwt: uses the HTTP request body to send `client_id`, `client_assertion_type`, and `client_assertion` as
// `application/x-www-form-urlencoded` body parameters. The `client_assertion` is signed using a private key supplied as an clientPrivateKey options parameter}.
//
// see: [RFC 6749 - The OAuth 2.0 Authorization Framework](https://www.rfc-editor.org/rfc/rfc6749.html#section-2.3)
// see [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html#ClientAuthentication)
// see [OAuth Token Endpoint Authentication Methods](https://www.iana.org/assignments/oauth-parameters/oauth-parameters.xhtml#token-endpoint-auth-method)
type ClientAuthenticationMethodType string

//this is coupled with the TokenEndpointAuthMethodsSupported field in the Endpoint Definition

const (
	ClientAuthenticationMethodTypeClientSecretBasic ClientAuthenticationMethodType = "client_secret_basic"
	ClientAuthenticationMethodTypePrivateKeyJwt     ClientAuthenticationMethodType = "private_key_jwt"
)
