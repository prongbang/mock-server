package repository

// DecryptRepository is interface
type DecryptRepository interface {
}

type decryptRepository struct {
}

// NewDecryptRepository is funcation create instance
func NewDecryptRepository() DecryptRepository {
	return &decryptRepository{}
}
