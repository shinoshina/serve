package netbase

import (
	"fmt"

	"golang.org/x/sys/unix"
	"goserve/internal/base/logger"
)

type engine struct {
	main_evl *eventloop

	sub_evls []*eventloop

	evl_counts map[*eventloop]int
	evl_map    map[int32](*eventloop)

	efd int

	num_loops int
	next_loop int
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

func (eng *engine) register(fd int32,num int) {

	eng.sub_evls[num].register(int(fd))
	eng.evl_map[fd] = eng.sub_evls[num]

}
func (eng *engine) accept(fd int32) (conn *connction) {

	clifd, cliaddr, _ := unix.Accept4(int(fd), unix.SOCK_NONBLOCK)

	conn = newConnection(int32(clifd))
	conn.raddr.Raw_address = cliaddr.(*unix.SockaddrInet4).Addr
	conn.raddr.Port = cliaddr.(*unix.SockaddrInet4).Port

	fmt.Printf("new connection\naddress : %v \nport : %v\nnumber %v loops listening\n", conn.raddr.Raw_address, conn.raddr.Port,eng.next_loop)

	eng.register(int32(clifd),eng.dispatch())
	return

}

func (e *engine) Launch() {

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0) // | unix.SOCK_NONBLOCK
	if err != nil {
		fmt.Println(err)
	}
	e.efd = fd

	fmt.Printf("engine fd : %v \n", e.efd)
	logger.Debugf("this is a debug message")
	

	address := [4]byte{127, 0, 0, 1}
	addr := &unix.SockaddrInet4{
		Port: 4211,
		Addr: address,
	}

	unix.Bind(fd, addr)
	unix.Listen(fd, 5)

	e.main_evl.register(fd)

}

func (eng *engine) Start(){

	for i:= 0;i<eng.num_loops;i++{
		go eng.sub_evls[i].loop()
	}
	eng.main_evl.loop()
}
