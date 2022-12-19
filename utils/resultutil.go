package utils

const (
	SUCCESS                    = 200
	ERROR                      = 500
	INVALID_PARAMS             = 400
	ValidationError            = 4010
	ValidationErrorMalformed   = 4011
	ValidationErrorExpired     = 4012
	ValidationErrorNotValidYet = 4013
	ValidationErrorAudience    = 4014
	ValidationErrorIssuer      = 4015
)

var MsgFlags = map[int]string{
	SUCCESS:                    "SUCCESS",
	ERROR:                      "ERROR",
	INVALID_PARAMS:             "INVALID_PARAMS",
	ValidationError:            "ValidationError",
	ValidationErrorMalformed:   "ValidationErrorMalformed",
	ValidationErrorExpired:     "ValidationErrorExpired",
	ValidationErrorNotValidYet: "ValidationErrorNotValidYet",
	ValidationErrorAudience:    "ValidationErrorAudience",
	ValidationErrorIssuer:      "ValidationErrorIssuer",
}

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SuccessMsg
func Success(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: SUCCESS,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

// FailMsg
func Fail(code int) *ResponseMsg {
	msgObj := &ResponseMsg{
		Code: code,
		Msg:  GetMsg(code),
	}
	return msgObj
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
