package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const accountNotFoundError = "account not found"

func AccountNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", accountNotFoundError, id))
}

func IsAccountNotFound(err error) bool {
	return strings.Contains(err.Error(), accountNotFoundError)
}
