package netbase

import "fmt"

type Connection interface {
	read() (int, error)
	write()
}
type Engine struct {
	E *engine
}
type (
	EventHandler interface {
		onMessageArrival(c Connection)

		onConnect()

		onDisconnect()
	}

	DefaultHandler struct{}
)

func (d *DefaultHandler) onConnect() {

	fmt.Println("handler print : connection")

}
func (d *DefaultHandler) onMessageArrival(c Connection) {

	fmt.Println("handler print : message")

}

func (d *DefaultHandler) onDisconnect() {
	fmt.Println("handler print : disconnect")

}
