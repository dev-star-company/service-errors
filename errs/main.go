package errs

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RequesterIDRequiredError = "requester id is required"
	invalidForeignKeyError   = "invalid foreign key"
	listingError             = "error listing"
	startTransactionError    = "starting transaction"
	commitTransactionError   = "commiting transaction"
	deleteError              = "error deleting"
	saveError                = "error saving"
	createError              = "error creating"
	invalidOrderByValue      = "invalid order by value"
	badRequestError          = "bad request"
	unknownError             = "unknown error occurred"
	internalErrorMessage     = "internal error occurred"
)

func UnknownError(err error) error {
	return status.Errorf(codes.Unknown, "%s: %v", unknownError, err)
}

func IsUnknownError(err error) bool {
	return strings.Contains(err.Error(), unknownError) || status.Code(err) == codes.Unknown
}

func RequesterIDRequired() error {
	return errors.New(RequesterIDRequiredError)
}

func IsRequesterIDRequired(err error) bool {
	return strings.Contains(err.Error(), RequesterIDRequiredError)
}

func CreateError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", createError, err)
}

func IsCreateError(err error) bool {
	return strings.Contains(err.Error(), createError)
}

func SavingError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", saveError, err)
}

func IsSavingError(err error) bool {
	return strings.Contains(err.Error(), saveError)
}

func DeleteError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", deleteError, err)
}

func IsDeleteError(err error) bool {
	return strings.Contains(err.Error(), deleteError)
}

func CommitTransactionError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", commitTransactionError, err))
}

func IsCommitTransactionError(err error) bool {
	return strings.Contains(err.Error(), commitTransactionError)
}

func StartTransactionError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", startTransactionError, err))
}

func IsStartTransactionError(err error) bool {
	return strings.Contains(err.Error(), startTransactionError)
}

func ListingError(entity string, err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", listingError, err))
}

func IsListingError(err error) bool {
	return strings.Contains(err.Error(), listingError)
}

func InvalidForeignKey(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", invalidForeignKeyError, err))
}

func IsInvalidForeignKey(err error) bool {
	return strings.Contains(err.Error(), invalidForeignKeyError)
}

func InvalidOrderByValue(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", invalidOrderByValue, err))
}

func IsInvalidOrderByValue(err error) bool {
	return strings.Contains(err.Error(), invalidOrderByValue)
}

func BadRequest(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", badRequestError, err))
}

func IsBadRequest(err error) bool {
	return strings.Contains(err.Error(), badRequestError)
}

func InternalError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", internalErrorMessage, err))
}

func IsInternalError(err error) bool {
	return strings.Contains(err.Error(), internalErrorMessage) || status.Code(err) == codes.Internal
}
