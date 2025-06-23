package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const media_typesNotFoundError = "media_types not found"

func MediaTypesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", media_typesNotFoundError, id))
}

func IsMediaTypesNotFound(err error) bool {
	return strings.Contains(err.Error(), media_typesNotFoundError)
}
