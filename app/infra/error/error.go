package error_wrapper

type ErrorWrapper struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (err ErrorWrapper) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}

func (err ErrorWrapper) Unwrap() error {
	return err.Err
}
func NewErrorWrapper(code int, err error, message string) error {
	return ErrorWrapper{
		Message: message,
		Code:    code,
		Err:     err,
	}
}
