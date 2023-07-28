package exception

import (
	"backend/graph/model"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ExceptionHandler(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code":   "401",
			"status": fmt.Sprintf("%d", http.StatusBadRequest),
		},
	}
}

type ErrorType string

const (
	NOT_FOUND        ErrorType = "NOT_FOUND"
	AUTHORIZATION    ErrorType = "AUTHORIZATION"
	BAD_REQUEST      ErrorType = "BAD_REQUEST"
	SERVER_ERROR     ErrorType = "SERVER_ERROR"
	VALIDATION_ERROR ErrorType = "VALIDATION_ERROR"
)

func MutationErrorHandler(ctx context.Context, err error, errorType ErrorType, field *string) model.MutationError {
	switch errorType {
	case NOT_FOUND:
		return model.NotFoundError{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
	case AUTHORIZATION:
		return model.AuthorizationError{
			Message: err.Error(),
			Code:    http.StatusUnauthorized,
		}
	case BAD_REQUEST:
		return model.BadRequestError{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
	case SERVER_ERROR:
		return model.ServerError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	case VALIDATION_ERROR:
		return model.CreateValidationError(err.Error(), *field)
	default:
		return nil
	}
}

func QueryErrorHandler(ctx context.Context, err error, errorType ErrorType, message *string) model.QueryError {
	switch errorType {
	case NOT_FOUND:
		return model.NotFoundError{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
	case AUTHORIZATION:
		return model.AuthorizationError{
			Message: err.Error(),
			Code:    http.StatusUnauthorized,
		}
	case BAD_REQUEST:
		return model.BadRequestError{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
	case SERVER_ERROR:
		return model.ServerError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	default:
		return nil
	}
}

func DatabaseError(err error, message string) error {
	if errors.As(err, &UniqueConstraintError{}) {
		return fmt.Errorf("%s should be unique", message)
	}
	return nil
}
