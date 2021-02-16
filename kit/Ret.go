package kit

type Ret struct {
	Code int 		 `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = "Success"
	FAIL = "Fail"
)

func Ok() (r Ret) {
	r.Code = 200
	r.Msg = SUCCESS
	return r
}

func Fail() (r Ret) {
	r.Code = -1
	r.Msg = FAIL
	return r
}

func FailAndMsg(msg string) (r Ret) {
	r.Code = -1
	r.Msg = msg
	return r
}
func OkAndData(data interface{})   (r Ret) {
	r.Code = 200
	r.Msg = SUCCESS
	r.Data = data
	return r
}