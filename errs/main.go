package errs

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	requesterIdRequiredError = "requester id is required"
	invalidForeignKeyError   = "invalid foreign key"
	listingError             = "error listing"
	startTransactionError    = "starting transaction"
	commitTransactionError   = "commiting transaction"
	deleteError              = "error deleting"
	saveError                = "error saving"
	createError              = "error creating"
	invalidOrderByValue      = "invalid order by value"
)

func RequesterIdRequired() error {
	return errors.New(requesterIdRequiredError)
}

func CreateError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", createError, err)
}

func SavingError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", saveError, err)
}

func IsRequesterIdRequired(err error) bool {
	return strings.Contains(err.Error(), requesterIdRequiredError)
}

func DeleteError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", deleteError, err)
}

func CommitError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", commitTransactionError, err))
}

func StartError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", startTransactionError, err))
}

func ListingError(entity string, err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", listingError, err))
}

func InvalidForeignKey(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", invalidForeignKeyError, err))
}

func InvalidOrderByValue(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", invalidOrderByValue, err))
}
