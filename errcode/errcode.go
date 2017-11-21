package errcode

type errinfo struct {
	errMsg string
}

const (
	ErrOk      = 0
	ErrAe      = 1
	ErrNotImpl = 2
	ErrPm      = 3
	ErrNe      = 4
	ErrUnPw    = 5
	ErrExp     = 6
)

const (
	ErrStrOk      = "ok"
	ErrStrAe      = "already exist"
	ErrStrNotImpl = "not implemented"
	ErrStrPm      = "parameter missing"
	ErrStrNe      = "not exist"
	ErrStrUnPw    = "username or password wrong"
	ErrStrExp     = "expired"
)

type ErrCodeType int

var ErrCode = make(map[ErrCodeType]errinfo)

func init() {
	ErrCode[ErrOk] = errinfo{ErrStrOk}
	ErrCode[ErrAe] = errinfo{ErrStrAe}
	ErrCode[ErrNotImpl] = errinfo{ErrStrNotImpl}
	ErrCode[ErrPm] = errinfo{ErrStrPm}
	ErrCode[ErrNe] = errinfo{ErrStrNe}
	ErrCode[ErrUnPw] = errinfo{ErrStrUnPw}
	ErrCode[ErrExp] = errinfo{ErrStrExp}
}
