package netbase

import (
	"goserve/internal/base/logger"

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

	conn_map map[int32](*connction)

	epoll_handler epoll_callback

	epoller *poller

	evl_type string
}

func (evl *eventloop) unregister(fd int32) {

	evl.epoller.unregister(fd)
	unix.Close(int(fd))

}
func (evl *eventloop) ComputeHandler(fd int32, event uint32) {

	if event&unix.EPOLLOUT != 0 {
		logger.Debugf("WRITE EVENT")
	}
	if event&InEvents != 0 {
		evl.read(evl.conn_map[fd])
	}
}

func (evl *eventloop) write(c *connction) {

}
func (evl *eventloop) read(c *connction) {

	n, err := c.read()
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
		}
	}
	c.write()

}

func (evl *eventloop) AcceptHandler(fd int32, event uint32) {
	if event&InEvents != 0 {
		c := evl.eng_from.accept(fd)
		evl.conn_map[c.fd] = c
	}
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

func (evl *eventloop) register(c *connction) {

	evl.conn_map[c.fd] = c
	evl.epoller.register(int(c.fd))
}
