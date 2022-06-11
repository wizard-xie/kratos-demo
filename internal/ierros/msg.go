package ierros

var message map[ICode]string

const (
	// MsgStatusOK success msg.
	MsgSuccess string = "操作成功"
	// MsgBadRequest invalidate argument msg.
	MsgBadRequest string = "参数错误，请校对!"
	// MsgUnauthorized use unauthorized msg.
	MsgUnauthorized string = "登录失效，请重新登录"
	// MsgForbidden permission denied msg.
	MsgForbidden string = "您没有访问权限"
	// MsgNotFound not found resource msg.
	MsgNotFound string = "没有找到资源"
	// MsgConflict conflict msg during server processing.
	MsgConflict string = "修改发生冲突，请稍后再试"
	// MsgTooManyRequests the msg request is too mundane.
	MsgTooManyRequests string = "请求太频繁啦，先休息一下吧"
	// ClientClosed is non-standard http status msg,
	// which defined by nginx.
	MsgClientClosed string = "连接已断开"
	// MsgInternalServerError the msg of internal server error.
	MsgInternalServerError string = "服务器开小差啦,稍后再来试一试"
	// MsgNotImplemented the server does not support a feature required by the current request.
	MsgNotImplemented string = "暂未支持该功能"
	// MsgServiceUnavailable the msg of service unavailable.
	MsgServiceUnavailable string = "服务不可用，非常抱歉"
	// MsgTimeout the msg of time out.
	MsgTimeout string = "服务器处理超时啦，请稍后再试"
	// MsgUnknown msg of unknown code for message map.
	MsgUnknown = "未知错误"
)

func init() {
	message = make(map[ICode]string)
	message[CodeSuccess] = MsgSuccess
	message[CodeBadRequest] = MsgBadRequest
	message[CodeUnauthorized] = MsgUnauthorized
	message[CodeForbidden] = MsgForbidden
	message[CodeNotFound] = MsgNotFound
	message[CodeTooManyRequests] = MsgTooManyRequests
	message[CodeClientClosed] = MsgClientClosed
	message[CodeInternalServerError] = MsgInternalServerError
	message[CodeNotImplemented] = MsgNotImplemented
	message[CodeServiceUnavailable] = MsgServiceUnavailable
	message[CodeTimeout] = MsgTimeout
}

// CodeMsg get msg by code
func CodeMsg(code ICode) string {
	m, ok := message[code]
	if !ok {
		return MsgUnknown
	}

	return m
}
