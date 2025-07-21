package ports

type HashingPort interface {
	Hash(password string) (string, error)
	Compare(hashPassword, password string) bool
}
