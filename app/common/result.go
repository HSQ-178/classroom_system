package common

type Result struct {
	Code int         `json:"code"` //状态码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

var result Result

// Success 的成功返回
func Success(data interface{}) Result {
	result.Code = 200
	result.Msg = "success"
	result.Data = data
	return result
}

// Fail 失败
func Fail(msg string) Result {
	result.Code = 400
	result.Msg = msg
	return result
}

// FailWithCode 失败
func FailWithCode(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
	}
}

func Error() string {
	return result.Msg
}

// 构造函数
func NewResult() Result {
	return Result{}
}
