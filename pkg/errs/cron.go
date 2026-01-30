package errs

import "net/http"


var (
	ErrCronJobNameIsEmpty           = NewSystemError(SystemSubcategoryCron, 0, http.StatusInternalServerError, "cron job name is empty")
	ErrCronJobNotExistsOrNotEnabled = NewSystemError(SystemSubcategoryCron, 1, http.StatusInternalServerError, "cron job not exists or not enabled")
)
