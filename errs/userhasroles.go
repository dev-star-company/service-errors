package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userhasrolesNotFoundError = "user has roles not found"

func UserHasRolesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", userhasrolesNotFoundError, id))
}

func IsUserHasRolesNotFound(err error) bool {
	return strings.Contains(err.Error(), userhasrolesNotFoundError)
}
