package persistence

// DatabaseHandler is the interface that wraps the basic database operations.
type DatabaseHandler interface {
	CreateEvent(*Event) ([]byte, error)
	FindEvent(id string) (*Event, error)
	FindEventByName(name string) (*Event, error)
	FindAll() ([]Event, error)
}
