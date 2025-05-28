package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const accountTypeNotFoundError = "account type not found"

func AccountTypeNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", accountTypeNotFoundError, id))
}

func IsAccountTypeNotFound(err error) bool {
	return strings.Contains(err.Error(), accountTypeNotFoundError)
}
