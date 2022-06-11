package ierros

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type ICode int32

const (
	// CodeStatusOK success code.
	CodeSuccess ICode = 200
	// CodeBadRequest invalidate argument code.
	CodeBadRequest ICode = 400
	// CodeUnauthorized use unauthorized code.
	CodeUnauthorized ICode = 401
	// CodeForbidden permission denied code.
	CodeForbidden ICode = 403
	// CodeNotFound not found resource code.
	CodeNotFound ICode = 404
	// CodeConflict conflict code during server processing.
	CodeConflict ICode = 409
	// CodeTooManyRequests the code request is too mundane.
	CodeTooManyRequests ICode = 429
	// ClientClosed is non-standard http status code,
	// which defined by nginx.
	CodeClientClosed ICode = 499
	// CodeInternalServerError the code of internal server error.
	CodeInternalServerError ICode = 500
	// CodeNotImplemented the server does not support a feature required by the current request.
	CodeNotImplemented ICode = 501
	// CodeServiceUnavailable the code of service unavailable.
	CodeServiceUnavailable ICode = 503
	// CodeTimeout the code of time out.
	CodeTimeout ICode = 504
)

// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
func FromGRPCCode(code codes.Code) int32 {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return int32(CodeClientClosed)
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

// ToGRPCCode converts an error code to gRPC code.
// If the error code does not correspond to the GRPC status code,
// it will be directly converted to grpc code type.
func ToGRPCCode(code ICode) codes.Code {
	switch code {
	case CodeSuccess:
		return codes.OK
	case CodeBadRequest:
		return codes.InvalidArgument
	case CodeUnauthorized:
		return codes.Unauthenticated
	case CodeForbidden:
		return codes.PermissionDenied
	case CodeNotFound:
		return codes.NotFound
	case CodeConflict:
		return codes.Aborted
	case CodeTooManyRequests:
		return codes.ResourceExhausted
	case CodeInternalServerError:
		return codes.Internal
	case CodeNotImplemented:
		return codes.Unimplemented
	case CodeServiceUnavailable:
		return codes.Unavailable
	case CodeTimeout:
		return codes.DeadlineExceeded
	case CodeClientClosed:
		return codes.Canceled
	}

	return codes.Code(code)
}

// ToHTTPCode converts an  error code to HTTP status code.
// Error codes not declared in this package may not directly
// correspond to the http status code and return http.StatusInternalServerError directly.
func ToHTTPCode(code ICode) int {
	switch code {
	case CodeSuccess:
		return http.StatusOK
	case CodeBadRequest:
		return http.StatusBadGateway
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict:
		return http.StatusConflict
	case CodeTooManyRequests:
		return http.StatusTooManyRequests
	case CodeInternalServerError:
		return http.StatusInternalServerError
	case CodeNotImplemented:
		return http.StatusNotImplemented
	case CodeServiceUnavailable:
		return http.StatusServiceUnavailable
	case CodeTimeout:
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}
