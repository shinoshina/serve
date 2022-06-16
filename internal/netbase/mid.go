package netbase

type Engine struct {
	E *engine
}
type Connection interface {
	read() (int, error)
	write()
}
type EventHandler interface {
	onMessageArrival(c Connection)

	onConnect()

	onDisconnect()
}
