package snet

import (
	"fmt"

	"github.com/shinoshina/snet/internal/netbase"
)

type DefaultHandler struct{}

func (d *DefaultHandler) onConnect() {

	fmt.Println("handler print : connection")

}
func (d *DefaultHandler) onMessageArrival(c netbase.Connection) {

	fmt.Println("handler print : message")

}

func (d *DefaultHandler) onDisconnect() {
	fmt.Println("handler print : disconnect")

}

type Server struct {
	EventEngine *netbase.Engine

	Handler netbase.EventHandler
}

func DefaultEngine() (e *netbase.Engine) {

	e = new(netbase.Engine)
	e.E = netbase.NewEngine(3)
	return
}

func (s *Server) Launch() {

	s.buildInHandler()
	s.EventEngine.E.Launch()
	s.EventEngine.E.Start()

}

func (s *Server) buildInHandler() {
	s.EventEngine.E.BuiltInhandler(s.Handler)

}

// func main(){

// 	s := Server{
// 		EventEngine: DefaultEngine(),
// 		Handler: &netbase.DefaultHandler{},
// 	}
// 	s.Launch()
// }
