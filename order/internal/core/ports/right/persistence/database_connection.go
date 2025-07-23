package ports

type DatabaseConnectionPort interface {
	Connect() (interface{}, error)
	Close() error
}
