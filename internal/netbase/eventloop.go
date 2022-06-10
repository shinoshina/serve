package netbase

import (
	"fmt"

	"golang.org/x/sys/unix"
)

const (
	ErrEvents = unix.EPOLLERR | unix.EPOLLHUP | unix.EPOLLRDHUP

	OutEvents = ErrEvents | unix.EPOLLOUT

	InEvents = ErrEvents | unix.EPOLLIN | unix.EPOLLPRI
)

func (evl *eventloop) unregister(fd int32) {

	evl.epoller.unregister(fd)

}
func (evl *eventloop) ComputeHandler(fd int32, event uint32) {

	if event&InEvents != 0 {

		if event&unix.EPOLLRDHUP != 0 {

			evl.unregister(fd)
			unix.Close(int(fd))

		} else {

			p := make([]byte, 50)
			n, err := unix.Read(int(fd), p)
			if err != nil{
				fmt.Println(err)
			}
			if n > 0 {
				fmt.Println(p)
				fmt.Println(string(p))
				unix.Write(int(fd), p)
			}
		}

	}
}

func (evl *eventloop) AcceptHandler(fd int32, event uint32) {
	if event&InEvents != 0 {
		c := evl.eng_from.accept(fd)
		evl.conn_map[c.fd] = c
	}
}

type epoll_callback func(fd int32, event uint32)

type eventloop struct {
	eng_from *engine

	conn_map map[int32](*connction)

	epoll_handler epoll_callback

	epoller *poller

	evl_type string
}

func newEventloop(evl_type string) (evl *eventloop) {

	evl = new(eventloop)
	evl.evl_type = evl_type
	evl.conn_map = make(map[int32](*connction))

	if evl_type == "acceptor" {
		evl.epoll_handler = evl.AcceptHandler

	} else if evl_type == "compute" {
		evl.epoll_handler = evl.ComputeHandler
	}
	evl.epoller = newPoller()

	return

}
func (evl *eventloop) loop() {

	evl.epoller.Epoll(evl.epoll_handler)
}

func (evl *eventloop) register(fd int) {

	fmt.Printf("evl register fd : %v \n", fd)

	evl.epoller.register(fd)
}
