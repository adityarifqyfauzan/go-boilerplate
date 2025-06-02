// repository or data access layer implementation
package user

type LocalRepository interface {
}

type localRepository struct {
}

func NewRepository() LocalRepository {
	return &localRepository{}
}
