package persistence

type DatabaseHandler interface {
	CreateEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAll() ([]Event, error)
}
