package exception

type UnprocessableEntityException struct {
	Error string
}

func NewUnprocessableEntityException(s string) UnprocessableEntityException {
	return UnprocessableEntityException{Error: s}
}
