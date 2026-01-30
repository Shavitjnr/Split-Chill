package errs

import (
	"net/http"
)


var (
	ErrLoggingError = NewSystemError(SystemSubcategoryLogging, 0, http.StatusInternalServerError, "logging error")
)
