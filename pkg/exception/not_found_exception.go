package exception

type NotFoundException struct {
	Error string
}

func NewNotFoundException(s string) NotFoundException {
	return NotFoundException{Error: s}
}
