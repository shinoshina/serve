package test

import (
	"goserve/internal/netbase"
)

func main() {

	ac := netbase.Acceptor{}
	ac.Init_test()
	

	ep := netbase.NewEpoller()
	ep.AddChannel(ac.AcFd)

	evl := netbase.NewEventLoop()
	evl.Loop()

}
