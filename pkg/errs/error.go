package errs

import (
	"strings"
)


type ErrorCategory int32


const (
	CATEGORY_SYSTEM ErrorCategory = 1
	CATEGORY_NORMAL ErrorCategory = 2
)


const (
	SystemSubcategoryDefault  = 0
	SystemSubcategorySetting  = 1
	SystemSubcategoryDatabase = 2
	SystemSubcategoryMail     = 3
	SystemSubcategoryLogging  = 4
	SystemSubcategoryCron     = 5
)


const (
	NormalSubcategoryGlobal                 = 0
	NormalSubcategoryUser                   = 1
	NormalSubcategoryToken                  = 2
	NormalSubcategoryTwofactor              = 3
	NormalSubcategoryAccount                = 4
	NormalSubcategoryTransaction            = 5
	NormalSubcategoryCategory               = 6
	NormalSubcategoryTag                    = 7
	NormalSubcategoryDataManagement         = 8
	NormalSubcategoryMapProxy               = 9
	NormalSubcategoryTemplate               = 10
	NormalSubcategoryPicture                = 11
	NormalSubcategoryConverter              = 12
	NormalSubcategoryUserCustomExchangeRate = 13
	NormalSubcategoryModelContextProtocol   = 14
	NormalSubcategoryLargeLanguageModel     = 15
	NormalSubcategoryUserExternalAuth       = 16
	NormalSubcategoryOAuth2                 = 17
	NormalSubcategoryInsightsExplorer       = 18
	NormalSubcategoryTagGroup               = 19
)


type Error struct {
	Category       ErrorCategory
	SubCategory    int32
	Index          int32
	HttpStatusCode int
	Message        string
	BaseError      []error
	Context        any
}

type MultiErrors struct {
	errors []error
}


func (err *Error) Error() string {
	return err.Message
}


func (err *Error) Code() int32 {
	return int32(err.Category)*100000 + err.SubCategory*1000 + err.Index
}


func New(category ErrorCategory, subCategory int32, index int32, httpStatusCode int, message string, baseError ...error) *Error {
	return &Error{
		Category:       category,
		SubCategory:    subCategory,
		Index:          index,
		HttpStatusCode: httpStatusCode,
		Message:        message,
		BaseError:      baseError,
	}
}


func (err *MultiErrors) Error() string {
	if len(err.errors) == 1 {
		return err.errors[0].Error()
	}

	var ret strings.Builder
	var lastErrorChar byte

	ret.WriteString("multi errors: ")

	for i := 0; i < len(err.errors); i++ {
		if i > 0 {
			if lastErrorChar == '.' {
				ret.WriteString(" ")
			} else {
				ret.WriteString(", ")
			}
		}

		errorContent := err.errors[i].Error()
		lastErrorChar = errorContent[len(errorContent)-1]
		ret.WriteString(errorContent)
	}

	return ret.String()
}


func NewSystemError(subCategory int32, index int32, httpStatusCode int, message string) *Error {
	return New(CATEGORY_SYSTEM, subCategory, index, httpStatusCode, message)
}


func NewNormalError(subCategory int32, index int32, httpStatusCode int, message string) *Error {
	return New(CATEGORY_NORMAL, subCategory, index, httpStatusCode, message)
}


func NewLoggingError(message string, err ...error) *Error {
	return New(ErrLoggingError.Category,
		ErrLoggingError.SubCategory,
		ErrLoggingError.Index,
		ErrLoggingError.HttpStatusCode,
		message, err...)
}


func NewIncompleteOrIncorrectSubmissionError(err error) *Error {
	return New(ErrIncompleteOrIncorrectSubmission.Category,
		ErrIncompleteOrIncorrectSubmission.SubCategory,
		ErrIncompleteOrIncorrectSubmission.Index,
		ErrIncompleteOrIncorrectSubmission.HttpStatusCode,
		ErrIncompleteOrIncorrectSubmission.Message, err)
}


func NewErrorWithContext(baseError *Error, context any) *Error {
	return &Error{
		Category:       baseError.Category,
		SubCategory:    baseError.SubCategory,
		Index:          baseError.Index,
		HttpStatusCode: baseError.HttpStatusCode,
		Message:        baseError.Message,
		BaseError:      baseError.BaseError,
		Context:        context,
	}
}


func NewMultiErrorOrNil(errors ...error) error {
	count := len(errors)

	if count < 1 {
		return nil
	} else if count == 1 {
		return errors[0]
	}

	return &MultiErrors{
		errors: errors,
	}
}


func Or(err error, defaultErr *Error) *Error {
	if finalError, ok := err.(*Error); ok {
		return finalError
	} else {
		return defaultErr
	}
}


func IsCustomError(err error) bool {
	_, ok := err.(*Error)
	return ok
}
