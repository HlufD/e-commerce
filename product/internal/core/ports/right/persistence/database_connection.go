package persistence

type DatabaseConnectionPort interface {
	Connect() (interface{}, error)
	Close() error
}
