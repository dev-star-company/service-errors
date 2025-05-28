package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const currencyNotFoundError = "currency not found"

func CurrencyNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", currencyNotFoundError, id))
}

func IsCurrencyNotFound(err error) bool {
	return strings.Contains(err.Error(), currencyNotFoundError)
}
