package exception

type CredentialException struct {
	Error string
}

func NewCredentialException(s string) CredentialException {
	return CredentialException{Error: s}
}
