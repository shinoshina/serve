package netbase

type Engine struct {
	E *engine
}
type Connection interface {
	read() (int, error)
	write()
}
type EventHandler interface {
	MessageArrival(c Connection)

	Connect()

	Disconnect()
}
