package main

import "goserve/internal"

func main(){


	ac := internal.Acceptor{

	}
	ac.Init_test()


	ep := internal.NewEpoller()
	ep.AddChannel(ac.AcFd)


	evl := internal.NewEventLoop(ep)
	evl.Loop()



}