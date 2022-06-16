package netbase

import (
	"fmt"
	"github.com/shinoshina/snet/internal/base/logger"

	"golang.org/x/sys/unix"
)

type engine struct {
	main_evl *eventloop

	sub_evls []*eventloop

	evl_counts map[*eventloop]int
	evl_map    map[int32](*eventloop)

	efd int

	num_loops int
	next_loop int

	handler EventHandler
}

func NewEngine(num_loops int) (eng *engine) {

	eng = new(engine)

	eng.num_loops = num_loops
	eng.next_loop = 0

	eng.main_evl = newEventloop("acceptor")
	eng.main_evl.eng_from = eng

	eng.sub_evls = make([]*eventloop, num_loops)
	for i := 0; i < num_loops; i++ {
		eng.sub_evls[i] = newEventloop("compute")
	}

	eng.evl_map = make(map[int32]*eventloop)

	eng.evl_counts = make(map[*eventloop]int)
	for i := 0; i < num_loops; i++ {
		eng.evl_counts[eng.sub_evls[i]] = 0
	}

	return
}

func (eng *engine) dispatch()(num int){

	//round-robin
	num = eng.next_loop
	eng.next_loop++
	if eng.next_loop == eng.num_loops {
		eng.next_loop = 0
	}

    return
}

func (eng *engine) register(c *connection,num int) {

	eng.sub_evls[num].register(c)
	eng.evl_map[c.fd] = eng.sub_evls[num]

}
func (eng *engine) accept(fd int32) (c *connection) {

	clifd, cliaddr, _ := unix.Accept4(int(fd), unix.SOCK_NONBLOCK)

	c = newConnection(int32(clifd))
	c.local_addr.Address = cliaddr.(*unix.SockaddrInet4).Addr
	c.local_addr.Port = cliaddr.(*unix.SockaddrInet4).Port

	logger.Infof("new connection\naddress : %v \nport : %v\nnumber %v loops listening", c.local_addr.Address, c.local_addr.Port,eng.next_loop)

	eng.register(c,eng.dispatch())
	return

}

func (e *engine) Launch() {

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0) // | unix.SOCK_NONBLOCK
	if err != nil {
		fmt.Println(err)
	}
	e.efd = fd

	logger.Infof("engine fd : %v ", e.efd)

	addressb := [4]byte{127, 0, 0, 1}
	addr := &unix.SockaddrInet4{
		Port: 4211,
		Addr: addressb,
	}

	ac := newConnection(int32(fd))
	ac.local_addr.Address = addressb
	ac.local_addr.Port = 4211

	unix.Bind(fd, addr)
	unix.Listen(fd, 5)

	e.main_evl.register(ac)

}

func (eng *engine) Start(){

	for i:= 0;i<eng.num_loops;i++{
		go eng.sub_evls[i].loop()
	}
	eng.main_evl.loop()
}

func (eng *engine) BuiltInhandler(ev EventHandler){

	eng.handler = ev

	eng.main_evl.handler = ev
	for i := 0; i < eng.num_loops; i++ {
		eng.sub_evls[i].handler = ev
	}
}