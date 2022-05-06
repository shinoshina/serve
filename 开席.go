package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fd)
	address := [4]byte{127, 0, 0, 1}
	addr := &unix.SockaddrInet4{
		Port: 4201,
		Addr: address,
	}

	unix.Bind(fd, addr)


	unix.Listen(fd,5)

	for {
		clifd, cliaddr, err := unix.Accept(fd)

		if clifd != -1 {
			fmt.Println(clifd)
			fmt.Println(cliaddr)
		}else{
			fmt.Println(err)
		}
	}

}
