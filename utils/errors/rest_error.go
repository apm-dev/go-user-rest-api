package errors

import (
	"fmt"
	"github.com/apm-dev/go-user-rest-api/utils/logger"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var errorMessages = map[string]string{
	"required": "%s field is required",
	"email":    "%s field should be valid email address",
	"min":      "%s field does not provide min value",
	"max":      "%s field exceed the max value",
	"eqfield":  "%s field should be same as another field",
	"alpha":    "%s field should be alphabetical value",
}

type RestError struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Content interface{} `json:"content"`
}

func NotFound(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "NOT_FOUND",
	}
}

func NotSaved(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotAcceptable,
		Error:   "NOT_SAVED",
	}
}

func BadRequest(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BAD_REQUEST",
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_ERROR",
	}
}

func DatabaseError(err error, log string) *RestError {
	logger.Error(log, err)
	return InternalServerError("database error")
}

var (
	validationErrorPruneRegexp = regexp.MustCompile(`([a-zA-Z\W]+):([a-zA-Z0-9\d\n\S\W"'.]*):`)
	validationErrorKeysRegexp  = regexp.MustCompile(`'([a-zA-Z0-9]+)'`)
)

func ValidationError(obj interface{}, err string) *RestError {
	if !strings.Contains(err, "Field validation") {
		return &RestError{
			Message: "Request body should be a valid json",
			Status:  http.StatusBadRequest,
			Error:   "INVALID_INPUT",
			Content: nil,
		}
	}
	errs := strings.Split(err, "\n")
	errors := make(map[string]string, len(errs))
	for _, s := range errs {
		tmpErr := validationErrorPruneRegexp.ReplaceAllString(s, "")
		keys := validationErrorKeysRegexp.FindAllString(tmpErr, 2)
		fmt.Println("keys", keys, "len", len(keys))
		keys[0] = strings.Replace(keys[0], "'", "", -1)
		keys[1] = strings.Replace(keys[1], "'", "", -1)
		field, _ := reflect.TypeOf(obj).FieldByName(keys[0])
		key := field.Tag.Get("json")
		errors[key] = fmt.Sprintf(errorMessages[keys[1]], key)
	}
	return &RestError{
		Message: "Input validation error",
		Status:  http.StatusUnprocessableEntity,
		Error:   "INVALID_INPUT",
		Content: errors,
	}
}
