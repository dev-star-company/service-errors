package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	roleNotFoundError = "user not found"
)

func RoleNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", roleNotFoundError, id))
}

func IsRoleNotFound(err error) bool {
	return strings.Contains(err.Error(), roleNotFoundError)
}
