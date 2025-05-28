package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const transactionNotFoundError = "transaction not found"

func TransactionNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", transactionNotFoundError, id))
}

func IsTransactionNotFound(err error) bool {
	return strings.Contains(err.Error(), transactionNotFoundError)
}
