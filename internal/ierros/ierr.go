package ierros

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// IErrorHTTPResult the information that IError will expose to the client.
type IErrorHTTPResult struct {
	Code int
	Msg  string
}

type IError struct {
	// Code is error code.
	Code ICode `json:"code"`
	// FMsg error message displayed to client
	FMsg string `json:"f_msg"`
	// BMsg the backend needs to print the error information of the log.
	BMsg string `json:"b_msg"`
	// Level the level of IError.
	// this field can be passed in directly when calling the kratos log library.
	ILevel log.Level `json:"i_level"`
}

// Error convert IError to json string.
func (ie IError) Error() string {
	tpl := `{"code":%d,"f_msg":"%s","b_msg":"%s","i_level":%d}`

	return fmt.Sprintf(tpl, ie.Code, ie.FMsg, ie.BMsg, ie.ILevel)
}

// ToGRPCError convert IError to GRPC error.
func (ie IError) ToGRPCError() error {
	return status.Error(ToGRPCCode(ie.Code), ie.Error())
}

// ToHttpErrorResult convert IError to http status code and error respose body.
func (ie IError) ToHttpErrorResult() (int, IErrorHTTPResult) {
	return ToHTTPCode(ie.Code), IErrorHTTPResult{Code: int(ie.Code), Msg: ie.FMsg}
}

// PrintLog print log information of IError
func (ie IError) PrintLog(logger log.Logger) {
	info := struct {
		Code ICode
		Err  string
	}{
		Code: ie.Code,
		Err:  ie.BMsg,
	}

	logger.Log(ie.ILevel, info)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func (ie IError) Wrapf(format string, opts ...interface{}) error {
	err := fmt.Sprintf(format, opts...)
	ie.BMsg = err

	return errors.Wrapf(ie, err)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func (ie IError) Warp(message string) error {
	return errors.Wrap(ie, message)
}

// GRPCErrorToIError revert a GRPC error converted from IError to IError.
func GRPCErrorToIError(err error) (*IError, bool) {
	if err == nil {
		return nil, false
	}

	stu, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	ie, er := ToIError(stu.Message())
	if er != nil {
		return nil, false
	}

	return ie, true
}

// ToIError attempt to restore IError object from a string.
func ToIError(msg string) (*IError, error) {
	var ierr IError

	if err := json.Unmarshal([]byte(msg), &ierr); err != nil {
		return nil, err
	}

	return &ierr, nil
}

func IsIError(err error) (*IError, bool) {
	if err == nil {
		return nil, false
	}

	causeErr := errors.Cause(err)

	ie, ok := causeErr.(IError)
	if ok {
		return &ie, ok
	}

	return nil, false
}

// KratosGRPCIError the middleware of handle IError To GRPC error.
func KratosGRPCIError() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			rep, err := handler(ctx, req)

			ie, ok := IsIError(err)
			if ok {
				return rep, ie.ToGRPCError()
			}

			return rep, err
		}
	}
}
