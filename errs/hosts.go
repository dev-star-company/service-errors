package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const hostsNotFoundError = "hosts not found"

func HostsNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", hostsNotFoundError, id))
}

func IsHostsNotFound(err error) bool {
	return strings.Contains(err.Error(), hostsNotFoundError)
}
