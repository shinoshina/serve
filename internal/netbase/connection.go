package netbase

import (
	"goserve/internal/base/buffer"

	"golang.org/x/sys/unix"
)

type connction struct {
	evl_from    *eventloop
	buf         *buffer.Buffer // use string temply
	fd          int32
	local_addr  address
	remote_addr address
}

func newConnection(fd int32) (c *connction) {

	c = new(connction)
	c.fd = fd
	c.buf = buffer.NewBuffer(512)

	return
}

func (c *connction) write() {
	n, err := unix.Write(int(c.fd), c.buf.Raw[:c.buf.E])

	if err != nil || n == 0 {
		return
	} else {
		if c.buf.S += n; c.buf.S != c.buf.E {
			// indicates the kernel buffer is full blocked ,need registered EPOLLOUT to the poller,waiting
			// for the kernel buffer ready writed
		} else {
			c.buf.S = 0
			c.buf.E = 0
		}

	}

}
func (c *connction) read() (n int, err error) {

	// unix.Read will cover the buffer
	n, err = unix.Read(int(c.fd), c.buf.Raw)
	if err != nil || n == 0 {
		return
	} else {
		c.buf.E = n
	}
	return
}
