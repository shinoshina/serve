package netbase

import (
	"fmt"

	"golang.org/x/sys/unix"
)

type engine struct {
	main_evl *eventloop

	sub_evls []*eventloop

	evl_map map[int32](*eventloop)

	efd int
}

func NewEngine() (eng *engine) {

	eng = new(engine)
	eng.main_evl = newEventloop("acceptor")
	eng.main_evl.eng_from = eng

	eng.sub_evls = make([]*eventloop, 1)
	eng.sub_evls[0] = newEventloop("compute")

	eng.evl_map = make(map[int32]*eventloop)

	return
}

func (eng *engine) register(fd int32){

	eng.sub_evls[0].register(int(fd))
    eng.evl_map[fd] = eng.sub_evls[0]

}
func (eng *engine) accept(fd int32) (conn *connction) {

	clifd, cliaddr, _ := unix.Accept(int(fd))

	conn = newConnection(int32(clifd))
	conn.raddr.Raw_address = cliaddr.(*unix.SockaddrInet4).Addr
	conn.raddr.Port = cliaddr.(*unix.SockaddrInet4).Port
    
	fmt.Printf("new connection\naddress : %v \nport : %v\n",conn.raddr.Raw_address,conn.raddr.Port)
	eng.register(int32(clifd))
	return

}

func(e *engine) Start(){

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM | unix.SOCK_NONBLOCK, 0)
	if err!= nil{
		fmt.Println(err)
	}
	e.efd = fd

	fmt.Printf("engine fd : %v \n",e.efd)

	address := [4]byte{127, 0, 0, 1}
	addr := &unix.SockaddrInet4{
		Port: 4211,
		Addr: address,
	}

	unix.Bind(fd, addr)
	unix.Listen(fd, 5)

	e.main_evl.register(fd)

	go e.sub_evls[0].loop()

	e.main_evl.loop()

}
