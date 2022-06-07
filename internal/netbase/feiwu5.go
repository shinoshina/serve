package netbase




var(


	defaultLoopNums = 5
)


type Option struct{

}
type Server struct{

	Address TcpAddress
	Option Option

	accptor *Acceptor              //  mainloop
	subLoops []*EventLoop
	connMap  map[TcpAddress]int    //  address : loop number

}




func Default()(sv *Server){

	sv = new(Server)
	sv.subLoops = make([]*EventLoop, defaultLoopNums)

	for i:=0;i<defaultLoopNums;i++{

		sv.subLoops[i] = NewEventLoop()
	}

	sv.accptor = NewAcceptor()

	sv.Address = TcpAddress{
		Address: "127.0.0.1",
		Port: 4650,
	}

	return
}

func (sv *Server)RegisterCallback(kd string,cb EventCallback){

	readCb = cb

}


func (sv *Server)Start(){


	sv.accptor.listen(sv.Address)

	for i:=0;i<defaultLoopNums;i++{
		go sv.subLoops[i].Loop()
	}

}