package serializer

type Response struct {
	Code  int               `json:"code"`
	Data  any               `json:"data"`
	Msg   string            `json:"msg"`
	Error map[string]string `json:"error,omitempty"`
}

func MessageResponse(message string) Response {
	return Response{
		Msg: message,
	}
}
func GeneralResponse(message string, input any) Response {
	return Response{
		Msg:  message,
		Data: input,
	}
}
func ErrRequestFormat(code int, errors map[string]string) Response {
	return Response{
		Code:  code,
		Error: errors,
	}

}

// func ErrResponse(code int, message string, err error) Response {
// 	var errMsg string
// 	if err != nil {
// 		msg := err.Error()
// 		errMsg = msg
// 	}
// 	return Response{
// 		Code:  code,
// 		Msg:   message,
// 		Error: errMsg,
// 	}
// }
