package kk

import (
	"fmt"

	"golang.org/x/sys/unix"
)





type Acceptor struct{


	AcFd int
	AcFdChannel Channel

}

func defaultConnectionCb()(){


	fmt.Println("new connection")
}


func (ac *Acceptor)Init_test()(){

	fd,err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	ac.AcFd = fd
	if err != nil{
		fmt.Println(err)
	}

	address := [4]byte{127, 0, 0, 1}
	addr := &unix.SockaddrInet4{
		Port: 4201,
		Addr: address,
	}

	unix.Bind(fd, addr)
	unix.Listen(fd,5)


}