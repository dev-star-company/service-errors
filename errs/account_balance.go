package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const accountBalanceNotFoundError = "account balance not found"

func AccountBalanceNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", accountBalanceNotFoundError, id))
}

func IsAccountBalanceNotFound(err error) bool {
	return strings.Contains(err.Error(), accountBalanceNotFoundError)
}
