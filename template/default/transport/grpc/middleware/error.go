package middleware

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/payfazz/go-apt/pkg/fazzcommon/httpError"
	"github.com/payfazz/go-errors"
	"github.com/payfazz/tango/template/default/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TranslateToGrpcError() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			out, err = f(ctx, in)

			if nil == err {
				return out, nil
			}

			cause := err
			message := fmt.Sprint("[GRPC-ERROR] ", err.Error())
			if ge, ok := err.(*errors.Error); ok {
				cause = ge.Cause
				message = fmt.Sprint("[GRPC-ERROR] ", errors.FormatWithDeep(ge, config.DEFAULT_FORMAT_DEEP))
			}

			log.Println(message)

			return out, buildError(cause)
		}
	}
}

func buildError(err error) error {
	message := getHttpErrorMessage(err)

	if httpError.IsNotFoundError(err) {
		return status.Error(codes.NotFound, message)
	} else if httpError.IsConflictError(err) {
		return status.Error(codes.AlreadyExists, message)
	} else if httpError.IsBadRequestError(err) || httpError.IsUnprocessableEntityError(err) {
		return status.Error(codes.InvalidArgument, message)
	} else if httpError.IsGatewayTimeoutError(err) {
		return status.Error(codes.DeadlineExceeded, message)
	} else if httpError.IsNotImplementedError(err) {
		return status.Error(codes.Unimplemented, message)
	} else if httpError.IsForbiddenError(err) || httpError.IsUnauthorizedError(err) {
		return status.Error(codes.Unauthenticated, message)
	} else {
		return status.Error(codes.Internal, message)
	}
}

func getHttpErrorMessage(err error) string {
	if herr, ok := err.(httpError.HttpErrorInterface); ok {
		return herr.GetDetail().(string)
	}

	return err.Error()
}
