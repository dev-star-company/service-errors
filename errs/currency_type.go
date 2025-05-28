package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const currencyTypeNotFoundError = "currencyType not found"

func CurrencyTypeNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", currencyTypeNotFoundError, id))
}

func IsCurrencyTypeNotFound(err error) bool {
	return strings.Contains(err.Error(), currencyTypeNotFoundError)
}
