package ierros

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	badRequest    = NewBadRequest()
	grpcBadReqest = badRequest.ToGRPCError()
)

func Test_New(t *testing.T) {
	code := ICode(50)
	msg := "test_new"
	leve := log.LevelInfo

	err := New(Code(code), FMsg(msg), LogLevel(leve))

	if err.Code != code || err.FMsg != msg || err.ILevel != leve {
		t.Fail()
	}
}

func Test_GRPCErrorToIError(t *testing.T) {
	err, ok := GRPCErrorToIError(grpcBadReqest)
	if !ok {
		t.Fail()
	}

	if err.Code != badRequest.Code {
		t.Error("err.code != badRequest.Code")
	}

}
