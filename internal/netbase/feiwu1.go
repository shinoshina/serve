
// already shit,bu yao le
package netbase

import (
	"fmt"

	"golang.org/x/sys/unix"
)


const(

	NormalChannel int = 0
	AcceptorChannel int = 1
)


//TO DO : GETTER SETTER FOR BETTER CODING STYLE
type EventCallback func(fd int) (state string)

var readCb EventCallback = func(fd int) string {

	clifd, cliaddr, err := unix.Accept(fd)

	if clifd != -1 {
		fmt.Println(clifd)
		fmt.Println(cliaddr)

		_, t := cliaddr.(*unix.SockaddrInet4)
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}

	fmt.Println("handle read event")
	return "1"
}
var writeCb EventCallback
var errCb EventCallback

///////////////////////////////////////////////////////////////////////////////////////
/******HERE******HERE******/ var disconnCb EventCallback /******HERE******HERE********/ 
///////////////////////////////////////////////////////////////////////////////////////
type Channel struct {
	fd     int
	revent int
	
}

func (ch *Channel) setRevent(event int) {
	ch.revent = event
}

func (ch *Channel) handleEvent() {

	if ch.revent&InEvents != 0 {

		fmt.Printf("channel fd : %v\n", ch.fd)

		readCb(ch.fd)

	}

	// }else if ch.revent & OutEvents != 0{

	// 	if ch.writeCb != nil{
	// 		ch.writeCb()
	// 	}
	// }
}
