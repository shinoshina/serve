package snet

import "fmt"

type (
	Connection interface {
		read() (int, error)

		write()
	}
	EventHandler interface {
		MessageArrival(c Connection)

		Connect()

		Disconnect()
	}
	DefaultHandler struct{}
)

func (d DefaultHandler) Connect() {

	fmt.Println("handler print : connection")

}
func (d DefaultHandler) MessageArrival(c Connection) {

	fmt.Println("handler print : message")

}

func (d DefaultHandler) Disconnect() {
	fmt.Println("handler print : disconnect")

}

type Server struct {
	EventEngine *engine

	Handler EventHandler
}

func DefaultEngine() (e *engine) {

	e = NewEngine(3)
	return
}

func (s *Server) Launch() {

	s.buildInHandler()
	s.EventEngine.launch()
	s.EventEngine.start()

}

func (s *Server) buildInHandler() {
	s.EventEngine.builtInhandler(s.Handler)

}

// func main(){

// 	s := Server{
// 		EventEngine: DefaultEngine(),
// 		Handler: DefaultHandler{},
// 	}
// 	s.Launch()
// }
