package serializer

type Response[T any] struct {
	Code  int    `json:"code"`
	Data  *T     `json:"data"`
	Msg   string `json:"msg"`
	Error string `json:"error,omitempty"`
}

func GeneralResponse[T any](code int, message string, input *T) Response[T] {
	return Response[T]{
		Code: code,
		Msg:  message,
		Data: input,
	}
}
