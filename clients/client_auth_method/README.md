# Client Authentication Methods

This package contains helper functions used by data-source clients to refresh and
introspect OAuth tokens for different authentication styles. Each flow wires the
appropriate refresh and introspection helpers together so that call-sites can
remain agnostic of the exact grant implementation.

## JWT Bearer clients

JWT bearer clients authenticate with a signed assertion generated from their
private key. They typically operate without a client secret and rely on the
JWT to assert their identity during both refresh and introspection calls.

* `PrivateKeyJWTBearerRefreshToken` (see `private_key_jwt.go`) constructs a new
  refresh token request where the client assertion is produced by signing the
  JWT with the configured private key. The helper injects the JWT-specific
  parameters (`client_assertion` and `client_assertion_type`) before invoking
  the token endpoint.
* `PrivateKeyJWTBearerIntrospectToken` (also in `private_key_jwt.go`) performs
  a similar assertion-based authentication against the introspection endpoint.
  It signs a short-lived JWT, attaches it to the request, and then parses the
  server response.

When using the JWT bearer flow, both helper functions must be called so the
client reuses the same signing material for refresh and introspection.

## Confidential clients

Confidential clients rely on a traditional client id / client secret pair and
use the HTTP Basic auth header for token operations.

* `ClientSecretBasicRefreshToken` (see `client_secret_basic.go`) refreshes the
  token by adding the Basic-auth header computed from the configured client id
  and secret when calling the token endpoint.
* `ClientSecretBasicAuthIntrospectToken` (same file) invokes the introspection
  endpoint with the identical Basic-auth header so both requests rely on the
  shared secret.

This is the default confidential-client flow and should be used unless a vendor
has specific requirements.

## Epic (custom confidential client)

Epic systems expect confidential clients to refresh via client secret basic
but introspect using a bearer token.

* `ClientSecretBasicRefreshToken` handles the refresh the same way as the
  generic confidential-client flow.
* `ClientBearerTokenAuthIntrospectToken` (see `client_bearer_token.go`) injects
  the Epic-issued bearer token into the `Authorization` header when performing
  introspection instead of using Basic auth.

Use this combination only for Epic integrations; other vendors generally follow
one of the two flows above.
