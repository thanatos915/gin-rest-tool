package errSys

import "fmt"

type Error struct {
	// 错误码
	code int `json:"code"`
	// 错误消息
	message string `json:"message"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, message: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.message
}

func (e *Error) SetMsg(msg string) {
	e.message = msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.message, args...)
}
