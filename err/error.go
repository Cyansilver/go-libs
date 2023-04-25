package err

type Error struct {
	Code       int32
	Msg        string
	Err        error
	HttpStatus int
}

func New(code int32, msg string) *Error {
	httpStatus := 400
	if code >= 13 && code <= 23 {
		httpStatus = 400
	}
	if code == ERR_INVALID_TOKEN_CODE {
		httpStatus = 401
	}
	return &Error{
		Code:       code,
		Msg:        msg,
		HttpStatus: httpStatus,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) Unwrap() error {
	return e.Err
}
