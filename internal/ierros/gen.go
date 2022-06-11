package ierros

import "github.com/go-kratos/kratos/v2/log"

// Options parameters option declaration for creating IError.
type Options struct {
	Code     ICode
	FMsg     string
	LogLevel log.Level
}

//LogLevel is set the IError log level.
func LogLevel(level log.Level) Option {
	return func(o *Options) {
		o.LogLevel = level
	}
}

// FMsg is set the error message displayed on the front.
func FMsg(msg string) Option {
	return func(o *Options) {
		o.FMsg = msg
	}
}

// Code is set the code of IError.
func Code(code ICode) Option {
	return func(o *Options) {
		o.Code = code
	}
}

// Option is a func to fill Options fields.
type Option func(*Options)

// New an IError with options.
// If no option is set, it will return IError of internal server error
// and the error level is log.LevelError.
func New(opts ...Option) IError {
	opt := Options{
		LogLevel: log.LevelError,
		Code:     CodeInternalServerError,
		FMsg:     message[CodeInternalServerError],
	}

	for _, v := range opts {
		v(&opt)
	}

	return IError{
		Code:   opt.Code,
		FMsg:   opt.FMsg,
		ILevel: opt.LogLevel,
	}
}

// NewIErrorWithCodeMsg new an IError with code msg.
func NewIErrorWithCodeMsg(code int, msg string) IError {
	return newIErrorWithCodeMsg(ICode(code), msg)
}

func newIErrorWithCodeMsg(code ICode, msg string) IError {
	return New(Code(code), FMsg(msg))
}

// NewIErrorWithCodeMsgLevel new an IError with code msg and Ilevel
func NewIErrorWithCodeMsgLevel(code int, msg string, level log.Level) IError {
	return New(Code(ICode(code)), FMsg(msg), LogLevel(level))
}

// NewIErrorWithCode new an IError with code.
// if the code is not define in this packege, the IError.FMsg is ierros.MsgUnkonw.
func NewIErrorWithCode(code int) IError {
	return newIErrorWithCode(ICode(code))
}

func newIErrorWithCode(code ICode) IError {
	msg := message[code]

	return newIErrorWithCodeMsg(code, msg)
}

// NewBadRequest new bad request IError and the error level is log.LevelError.
func NewBadRequest() IError {
	return newIErrorWithCode(CodeBadRequest)
}

// NewUnauthorized new unauthorized IError and the error level is log.LevelError.
func NewUnauthorized() IError {
	return newIErrorWithCode(CodeUnauthorized)
}

// NewForbidden new forbidden IError and the error level is log.LevelError.
func NewForbidden() IError {
	return newIErrorWithCode(CodeForbidden)
}

// NewNotFound new not found IError and the error level is log.LevelError.
func NewNotFound() IError {
	return newIErrorWithCode(CodeNotFound)
}

// NewConflict new conflict IError and the error level is log.LevelError.
func NewConflict() IError {
	return newIErrorWithCode(CodeConflict)
}

// NewTooManyRequests new too too many request IError and the error level is log.LevelError.
func NewTooManyRequests() IError {
	return newIErrorWithCode(CodeTooManyRequests)
}

// NewClientClosed new client close IError and the error level is log.LevelError.
func NewClientClosed() IError {
	return newIErrorWithCode(CodeClientClosed)
}

// NewInternalServerError new internal server error IError and the error level is log.LevelError.
func NewInternalServerError() IError {
	return newIErrorWithCode(CodeInternalServerError)
}

// NewNotImplemented new not implemented IError and the error level is log.LevelError.
func NewNotImplemented() IError {
	return newIErrorWithCode(CodeNotImplemented)
}

// NewServiceUnavailable new service unavailable IError and the error level is log.LevelError.
func NewServiceUnavailable() IError {
	return newIErrorWithCode(CodeServiceUnavailable)
}

// NewTimeout new timeout IError and the error level is log.LevelError.
func NewTimeout() IError {
	return newIErrorWithCode(CodeTimeout)
}
