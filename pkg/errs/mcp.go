package errs

import "net/http"


var (
	ErrMCPServerNotEnabled = NewNormalError(NormalSubcategoryModelContextProtocol, 0, http.StatusBadRequest, "mcp server is not enabled")
)
