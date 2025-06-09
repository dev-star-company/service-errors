package errs

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	userNotFoundInLoginError = "user not found"
	wrongPasswordError       = "wrong password"
)

func UserNotFoundInLogin() error {
	return status.Error(codes.NotFound, userNotFoundInLoginError)
}

func IsUserNotFoundInLogin(err error) bool {
	return strings.Contains(err.Error(), userNotFoundInLoginError)
}

func WrongPassword() error {
	return status.Error(codes.Unauthenticated, wrongPasswordError)
}

func IsWrongPassword(err error) bool {
	return strings.Contains(err.Error(), wrongPasswordError)
}
