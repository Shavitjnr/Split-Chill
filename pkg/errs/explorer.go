package errs

import "net/http"


var (
	ErrInsightsExplorerIdInvalid   = NewNormalError(NormalSubcategoryInsightsExplorer, 0, http.StatusBadRequest, "explorer id is invalid")
	ErrInsightsExplorerNotFound    = NewNormalError(NormalSubcategoryInsightsExplorer, 1, http.StatusBadRequest, "explorer not found")
	ErrInsightsExplorerDataInvalid = NewNormalError(NormalSubcategoryInsightsExplorer, 2, http.StatusBadRequest, "explorer data is invalid")
)
