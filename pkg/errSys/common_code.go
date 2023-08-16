package errSys

var (
	Success         = NewError(0, "成功")
	ServerError     = NewError(1000001, "服务错误")
	InvalidParams   = NewError(1000002, "入参错误")
	InvalidSign     = NewError(1000003, "签名校验错误")
	ServerRepair    = NewError(1000004, "系统维护中")
	TooManyRequests = NewError(1000005, "请求次数过多")
	ServerBusy      = NewError(1000006, "系统繁忙，请稍后再试")
)
