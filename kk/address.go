package kk

import (
	"strconv"
	"strings"
)

//import "golang.org/x/sys/unix"

type TcpAddress struct {
	Port    int
	Address string
}

func Convert(addr string) (raw_address [4]byte) {
	s := strings.SplitN(addr,".",4)
	//raw_address = make([]byte)
	for i := 0; i < 4; i++ {
		b,_ := strconv.Atoi(s[i])
		raw_address[i] = byte(b)
	}
	return raw_address

}
