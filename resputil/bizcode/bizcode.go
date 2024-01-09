package bizcode

type Code int

func (code Code) String() string {
	return codeText[code]
}

const (
	OK                  Code = 1001
	NeedLogin           Code = 1002
	InternalServerError Code = 5000
)

var codeText = map[Code]string{
	OK:                  "ok",
	InternalServerError: "内部错误",
	NeedLogin:           "未登陆",
}
