package errs

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	invalidToken   = "invalid token"
	malformedToken = "malformed token validation"
)

func InvalidToken() error {
	return status.Error(codes.Unauthenticated, invalidToken)
}

func IsInvalidToken(err error) bool {
	return strings.Contains(err.Error(), invalidToken)
}

func MalformedToken() error {
	return status.Error(codes.Unauthenticated, malformedToken)
}

func IsMalformedToken(err error) bool {
	return strings.Contains(err.Error(), malformedToken)
}
