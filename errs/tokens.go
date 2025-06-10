package errs

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	invalidToken   = "invalid token"
	malformedToken = "malformed token validation"
	expiredToken   = "expired token"
)

func InvalidToken(details string) error {
	msg := invalidToken
	if details != "" {
		msg += ": " + details
	}
	return status.Error(codes.Unauthenticated, msg)
}

func IsInvalidToken(err error) bool {
	return strings.Contains(err.Error(), invalidToken)
}

func MalformedToken(details string) error {
	msg := malformedToken
	if details != "" {
		msg += ": " + details
	}
	return status.Error(codes.Unauthenticated, msg)
}

func IsMalformedToken(err error) bool {
	return strings.Contains(err.Error(), malformedToken)
}

func ExpiredToken() error {
	return status.Error(codes.Unauthenticated, expiredToken)
}

func IsExpiredToken(err error) bool {
	return strings.Contains(err.Error(), expiredToken)
}
