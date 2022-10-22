package persistence

type DatabaseHandler interface {
	CreateEvent(Event) ([]byte, error)
	FindEvent(id []byte) (*Event, error)
	FindEventByName(name string) (*Event, error)
	FindAll() ([]Event, error)
}
