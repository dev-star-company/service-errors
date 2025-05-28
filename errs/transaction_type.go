package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const transactionTypeNotFoundError = "transaction type not found"

func TransactionTypeNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", transactionTypeNotFoundError, id))
}

func IsTransactionTypeNotFound(err error) bool {
	return strings.Contains(err.Error(), transactionTypeNotFoundError)
}
