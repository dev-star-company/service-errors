package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const bankNotFoundError = "bank not found"

func BankNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", bankNotFoundError, id))
}

func IsBankNotFound(err error) bool {
	return strings.Contains(err.Error(), bankNotFoundError)
}
