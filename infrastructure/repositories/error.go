package repositories

type ErrorConcurrentUpdating struct {
	error
}

func NewErrorConcurrentUpdating(err error) ErrorConcurrentUpdating {
	return ErrorConcurrentUpdating{err}
}
