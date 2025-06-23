package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const tagsNotFoundError = "tags not found"

func TagsNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", tagsNotFoundError, id))
}

func IsTagsNotFound(err error) bool {
	return strings.Contains(err.Error(), tagsNotFoundError)
}
