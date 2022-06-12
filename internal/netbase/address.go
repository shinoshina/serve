package netbase

import (
	"strconv"
	"strings"
)

//import "golang.org/x/sys/unix"

// type address struct {
// 	Port    int
// 	Address string
// }
type address struct{
	Port int
	Address [4]byte
}

func Convert(addr string) (ra [4]byte) {
	s := strings.SplitN(addr,".",4)
	//raw_address = make([]byte)
	for i := 0; i < 4; i++ {
		b,_ := strconv.Atoi(s[i])
		ra[i] = byte(b)
	}
	return

}
