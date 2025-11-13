package reqidmw

import (
	"net/http"

	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/common/pkg/errors"
)

const (
	InvalidRequestIdFormatError = "INVALID_REQUEST_ID_FORMAT_ERROR"
)

func NewInvalidRequestIdFormatError() *errors.APIError {
	return errors.NewAPIError(
		http.StatusBadRequest,
		InvalidRequestIdFormatError,
		"Request id is in invalid format.",
	)
}
