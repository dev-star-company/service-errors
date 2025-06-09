package errs

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userNotFoundInLoginError = "user not found"

func UserNotFoundInLogin(id int) error {
	return status.Error(codes.NotFound, userNotFoundInLoginError)
}

func IsUserNotFoundInLogin(err error) bool {
	return strings.Contains(err.Error(), userNotFoundInLoginError)
}
