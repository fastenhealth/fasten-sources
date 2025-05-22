package pkg

/**
 * Supported Client Authentication Methods.
 *
 * - **`client_secret_basic`** (default) uses the HTTP `Basic` authentication scheme to send
 *   {@link Client.client_id `client_id`} and {@link Client.client_secret `client_secret`} in an
 *   `Authorization` HTTP Header.
 * - **`client_secret_post`** uses the HTTP request body to send {@link Client.client_id `client_id`}
 *   and {@link Client.client_secret `client_secret`} as `application/x-www-form-urlencoded` body
 *   parameters.
 * - **`private_key_jwt`** uses the HTTP request body to send {@link Client.client_id `client_id`},
 *   `client_assertion_type`, and `client_assertion` as `application/x-www-form-urlencoded` body
 *   parameters. The `client_assertion` is signed using a private key supplied as an
 *   {@link AuthenticatedRequestOptions.clientPrivateKey options parameter}.
 * - **`none`** (public client) uses the HTTP request body to send only
 *   {@link Client.client_id `client_id`} as `application/x-www-form-urlencoded` body parameter.
 *
 * @see [RFC 6749 - The OAuth 2.0 Authorization Framework](https://www.rfc-editor.org/rfc/rfc6749.html#section-2.3)
 * @see [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html#ClientAuthentication)
 * @see [OAuth Token Endpoint Authentication Methods](https://www.iana.org/assignments/oauth-parameters/oauth-parameters.xhtml#token-endpoint-auth-method)
 */

type ClientAuthenticationMethodType string

const (
	ClientAuthenticationMethodTypeClientSecretBasic ClientAuthenticationMethodType = "client_secret_basic"
	ClientAuthenticationMethodTypePrivateKeyJwt     ClientAuthenticationMethodType = "private_key_jwt"
)
