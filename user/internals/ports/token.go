package ports

type TokenPort interface {
	Generate(id string) (string, error)
	Validate(token string) (string, error)
}
