package errs

import (
	"fmt"
	"net/http"
)


var (
	ErrIncompleteOrIncorrectSubmission = NewNormalError(NormalSubcategoryGlobal, 0, http.StatusBadRequest, "incomplete or incorrect submission")
	ErrOperationFailed                 = NewNormalError(NormalSubcategoryGlobal, 1, http.StatusInternalServerError, "operation failed")
	ErrRequestIdInvalid                = NewNormalError(NormalSubcategoryGlobal, 2, http.StatusInternalServerError, "request id is invalid")
	ErrCiphertextInvalid               = NewNormalError(NormalSubcategoryGlobal, 3, http.StatusInternalServerError, "ciphertext is invalid")
	ErrNothingWillBeUpdated            = NewNormalError(NormalSubcategoryGlobal, 4, http.StatusBadRequest, "nothing will be updated")
	ErrFailedToRequestRemoteApi        = NewNormalError(NormalSubcategoryGlobal, 5, http.StatusBadRequest, "failed to request third party api")
	ErrPageIndexInvalid                = NewNormalError(NormalSubcategoryGlobal, 6, http.StatusBadRequest, "page index is invalid")
	ErrPageCountInvalid                = NewNormalError(NormalSubcategoryGlobal, 7, http.StatusBadRequest, "page count is invalid")
	ErrClientTimezoneOffsetInvalid     = NewNormalError(NormalSubcategoryGlobal, 8, http.StatusBadRequest, "client timezone offset is invalid")
	ErrQueryItemsEmpty                 = NewNormalError(NormalSubcategoryGlobal, 9, http.StatusBadRequest, "query items cannot be blank")
	ErrQueryItemsTooMuch               = NewNormalError(NormalSubcategoryGlobal, 10, http.StatusBadRequest, "query items too much")
	ErrQueryItemsInvalid               = NewNormalError(NormalSubcategoryGlobal, 11, http.StatusBadRequest, "query items have invalid item")
	ErrParameterInvalid                = NewNormalError(NormalSubcategoryGlobal, 12, http.StatusBadRequest, "parameter invalid")
	ErrFormatInvalid                   = NewNormalError(NormalSubcategoryGlobal, 13, http.StatusBadRequest, "format invalid")
	ErrNumberInvalid                   = NewNormalError(NormalSubcategoryGlobal, 14, http.StatusBadRequest, "number invalid")
	ErrNoFilesUpload                   = NewNormalError(NormalSubcategoryGlobal, 15, http.StatusBadRequest, "no files uploaded")
	ErrUploadedFileEmpty               = NewNormalError(NormalSubcategoryGlobal, 16, http.StatusBadRequest, "uploaded file is empty")
	ErrExceedMaxUploadFileSize         = NewNormalError(NormalSubcategoryGlobal, 17, http.StatusBadRequest, "uploaded file size exceeds the maximum allowed size")
	ErrFailureCountLimitReached        = NewNormalError(NormalSubcategoryGlobal, 18, http.StatusBadRequest, "failure count exceeded maximum limit")
	ErrRepeatedRequest                 = NewNormalError(NormalSubcategoryGlobal, 19, http.StatusBadRequest, "repeated request")
	ErrIPForbidden                     = NewNormalError(NormalSubcategoryGlobal, 20, http.StatusBadRequest, "ip address is forbidden to access this resource")
)


func GetParameterInvalidMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid", field)
}


func GetParameterIsRequiredMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is required", field)
}


func GetParameterMustLessThanMessage(field string, param string) string {
	return fmt.Sprintf("parameter \"%s\" must be less than %s", field, param)
}


func GetParameterMustLessThanCharsMessage(field string, param string) string {
	return fmt.Sprintf("parameter \"%s\" must be less than %s characters", field, param)
}


func GetParameterMustMoreThanMessage(field string, param string) string {
	return fmt.Sprintf("parameter \"%s\" must be more than %s", field, param)
}


func GetParameterMustMoreThanCharsMessage(field string, param string) string {
	return fmt.Sprintf("parameter \"%s\" must be more than %s characters", field, param)
}


func GetParameterLengthNotEqualMessage(field string, param string) string {
	return fmt.Sprintf("parameter \"%s\" length is not equal to %s", field, param)
}


func GetParameterNotBeBlankMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" cannot be blank", field)
}


func GetParameterInvalidUsernameMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid username format", field)
}


func GetParameterInvalidEmailMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid email format", field)
}


func GetParameterInvalidCurrencyMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid currency", field)
}


func GetParameterInvalidHexRGBColorMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid color", field)
}


func GetParameterInvalidAmountFilterMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid amount filter", field)
}


func GetParameterInvalidTagFilterMessage(field string) string {
	return fmt.Sprintf("parameter \"%s\" is invalid tag filter", field)
}
