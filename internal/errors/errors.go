package errors

type ValidationErrors struct {
	Message string
}

func (v ValidationErrors) Error() string {
	return v.Message
}

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

type UnauthorizedError struct {
	Msg string
}

func (e *UnauthorizedError) Error() string {
	return e.Msg
}
