package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const proofNotFoundError = "proof not found"

func ProofNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", proofNotFoundError, id))
}

func IsProofNotFound(err error) bool {
	return strings.Contains(err.Error(), proofNotFoundError)
}
