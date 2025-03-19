package pkg

import "errors"

var (
	// ErrSMARTTokenRefreshFailure this error occurs when an error is encountered while trying to refresh the Access Token
	// This could be due to an expired/invalid refresh token, a client credentials error or a network error
	ErrSMARTTokenRefreshFailure = errors.New("ErrSMARTTokenRefreshFailure")

	// ErrResourcePatientFailure this error occurs when an error is encountered while trying to fetch a patient resource.
	// This should be a rare error, and is usually due to decoding errors, but it could be due to network errors as well.
	ErrResourcePatientFailure = errors.New("ErrResourcePatientFailure")

	// ErrResourceHttpError this error occurs when an error is encountered while trying to fetch a resource.
	// This could be due to a network error, a 404 error, a 403 or a 500 error.
	ErrResourceHttpError = errors.New("ErrResourceHttpError")

	// ErrResourceInvalidContent this error occurs when an error is encountered while trying to unmarshal a resource.
	ErrResourceInvalidContent = errors.New("ErrResourceInvalidContent")
)
