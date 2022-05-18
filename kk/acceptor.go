package kk

import (
	"fmt"

	"golang.org/x/sys/unix"
)

type Acceptor struct {
	AcFd        int
	AcFdChannel Channel
}

func defaultConnectionCb() {

	fmt.Println("new connection")
}

func NewAcceptor() (ac *Acceptor) {

	ac = new(Acceptor)
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
	}
	ac.AcFd = fd
	return

}

func (ac *Acceptor) Init_test() {

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	ac.AcFd = fd
	if err != nil {
		fmt.Println(err)
	}

	address := [4]byte{127, 0, 0, 1}
	addr := &unix.SockaddrInet4{
		Port: 4201,
		Addr: address,
	}

	unix.Bind(fd, addr)
	unix.Listen(fd, 5)

}


func (ac *Acceptor) listen(adr TcpAddress){


	addr := &unix.SockaddrInet4{
		Port: adr.Port,
		Addr: Convert(adr.Address),
	}

	unix.Bind(ac.AcFd,addr)
	unix.Listen(ac.AcFd,5)
}
