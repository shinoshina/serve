package snet

import (
	"github.com/shinoshina/snet/internal/base/logger"

	"golang.org/x/sys/unix"
)

const (
	ErrEvents = unix.EPOLLERR | unix.EPOLLHUP | unix.EPOLLRDHUP

	OutEvents = ErrEvents | unix.EPOLLOUT

	InEvents = ErrEvents | unix.EPOLLIN | unix.EPOLLPRI
)

type epoll_callback func(fd int32, event uint32)


type eventloop struct {
	eng_from *engine

	conn_map map[int32](*connection)

	epoll_handler epoll_callback

	epoller *poller

	evl_type string

	handler EventHandler
}

func (evl *eventloop) unregister(fd int32) {

	evl.epoller.unregister(fd)
	unix.Close(int(fd))

}
func (evl *eventloop) computeHandler(fd int32, event uint32) {

	if event&unix.EPOLLOUT != 0 {
		logger.Debugf("WRITE EVENT")
	}
	if event&InEvents != 0 {
		evl.read(evl.conn_map[fd])
	}
}

func (evl *eventloop) write(c *connection) {

}
func (evl *eventloop) read(c *connection) {

	n, err := c.Read()
	if err != nil || n == 0 {
		if err == unix.EAGAIN {
			logger.Errorf("%v", err)
		}
		if n == 0 {
			if err == unix.ECONNRESET {
				logger.Errorf("Connection shut down")
			}
			logger.Infof("Connection shut down")
			evl.unregister(c.fd)
			evl.handler.Disconnect()
		}
	}
	evl.handler.MessageArrival(c)

}

func (evl *eventloop) acceptHandler(fd int32, event uint32) {
	if event&InEvents != 0 {
		c := evl.eng_from.accept(fd)
		evl.conn_map[c.fd] = c
		evl.handler.Connect()
	}
}

func newEventloop(evl_type string) (evl *eventloop) {

	evl = new(eventloop)
	evl.evl_type = evl_type
	evl.conn_map = make(map[int32](*connection))

	if evl_type == "acceptor" {
		evl.epoll_handler = evl.acceptHandler

	} else if evl_type == "compute" {
		evl.epoll_handler = evl.computeHandler
	}
	evl.epoller = newPoller()

	return

}
func (evl *eventloop) loop() {

	evl.epoller.Epoll(evl.epoll_handler)
}

func (evl *eventloop) register(c *connection) {

	evl.conn_map[c.fd] = c
	evl.epoller.register(int(c.fd))
}
