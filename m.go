package main

import (
	"goserve/kk"
)

func main() {

	ac := kk.Acceptor{}
	ac.Init_test()
	

	ep := kk.NewEpoller()
	ep.AddChannel(ac.AcFd)

	evl := kk.NewEventLoop()
	evl.Loop()

}
