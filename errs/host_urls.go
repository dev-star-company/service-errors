package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const shostURLSNotFoundError = "shostURLS not found"

func HostURLSNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", shostURLSNotFoundError, id))
}

func IsHostURLSNotFound(err error) bool {
	return strings.Contains(err.Error(), shostURLSNotFoundError)
}
