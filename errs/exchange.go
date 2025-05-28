package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const exchangeNotFoundError = "exchange not found"

func ExchangeNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", exchangeNotFoundError, id))
}

func IsExchangeNotFound(err error) bool {
	return strings.Contains(err.Error(), exchangeNotFoundError)
}
