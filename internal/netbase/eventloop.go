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

func (evl *eventloop)ComputeHandler(fd int32, event uint32) {

	if event&InEvents != 0 {

	} else if event&OutEvents != 0 {

	}

}

func (evl* eventloop)AcceptHandler(fd int32,event uint32){


	if event & InEvents != 0 {

		evl.eng.accept(fd)

	}
}

type epoll_callback func(fd int32, event uint32)

type eventloop struct {

	eng *engine

	conn_map map[int32](*connction)

	epoll_handler epoll_callback

	epoller *poller

	evl_type string

}

func newEventloop(evl_type string) (evl *eventloop){

	evl = new(eventloop)

	evl.evl_type = evl_type

	evl.conn_map = make(map[int32](*connction))

	if evl_type == "acceptor" {
		evl.epoll_handler = evl.AcceptHandler


	}else if evl_type == "compute" {
		evl.epoll_handler = evl.ComputeHandler
	}

	evl.epoller = newPoller()

	return

}
func (evl *eventloop) loop() {

	evl.epoller.Epoll(evl.AcceptHandler)
}

func (evl *eventloop)register(fd int){

	fmt.Printf("evl register fd : %v \n",fd)

	evl.epoller.register(fd)
}
