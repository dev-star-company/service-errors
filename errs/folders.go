package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const foldersNotFoundError = "folders not found"

func FoldersNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", foldersNotFoundError, id))
}

func IsFoldersNotFound(err error) bool {
	return strings.Contains(err.Error(), foldersNotFoundError)
}
