package serializer

type Response struct {
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Msg   string `json:"msg"`
	Error string `json:"error,omitempty"`
}

func GeneralResponse(message string, input any) Response {
	return Response{
		Msg:  message,
		Data: input,
	}
}

func ErrResponse(code int, message string, err error) Response {
	var errMsg string
	if err != nil {
		msg := err.Error()
		errMsg = msg
	}
	return Response{
		Code:  code,
		Msg:   message,
		Error: errMsg,
	}
}
