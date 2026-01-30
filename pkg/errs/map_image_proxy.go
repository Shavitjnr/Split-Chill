package errs

import "net/http"


var (
	ErrMapProviderNotCurrent      = NewNormalError(NormalSubcategoryMapProxy, 0, http.StatusBadRequest, "specified map provider is not set")
	ErrImageExtensionNotSupported = NewNormalError(NormalSubcategoryMapProxy, 0, http.StatusNotFound, "specified image extension is not supported")
)
