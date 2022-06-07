package netbase

import (
	"fmt"

	"golang.org/x/sys/unix"
)

////////////////////////////////////////////////////////////
// type EpollEvent struct {
//
// 	Events uint32
// 	Fd int32
// 	Pad int32
//
// }
// EPOLLERR                                    = 0x8
// EPOLLET                                     = 0x80000000
// EPOLLEXCLUSIVE                              = 0x10000000
// EPOLLHUP                                    = 0x10
// EPOLLIN                                     = 0x1
// EPOLLMSG                                    = 0x400
// EPOLLONESHOT                                = 0x40000000
// EPOLLOUT                                    = 0x4
// EPOLLPRI                                    = 0x2
// EPOLLRDBAND                                 = 0x80
// EPOLLRDHUP                                  = 0x2000
// EPOLLRDNORM                                 = 0x40
// EPOLLWAKEUP                                 = 0x20000000
// EPOLLWRBAND                                 = 0x200
// EPOLLWRNORM                                 = 0x100
// EPOLL_CTL_ADD                               = 0x1
// EPOLL_CTL_DEL                               = 0x2
// EPOLL_CTL_MOD                               = 0x3
////////////////////////////////////////////////////////////
// const (
// 	ErrEvents = unix.EPOLLERR | unix.EPOLLHUP | unix.EPOLLRDHUP

// 	OutEvents = ErrEvents | unix.EPOLLOUT

// 	InEvents = ErrEvents | unix.EPOLLIN | unix.EPOLLPRI
// )

type Epoller struct {
	thisLoop    *EventLoop
	channelList []Channel
	sequence    uint16
	EpollFd     int
	eventList   []unix.EpollEvent
	eventNum    int
}

func NewEpoller() (ep *Epoller) {

	ep = new(Epoller)
	ep.channelList = make([]Channel, 0)
	ep.EpollFd = EpollCreate()
	ep.eventList = make([]unix.EpollEvent, 30)
	fmt.Printf("epoll fd : %v\n", ep.EpollFd)
	return ep
}

func EpollCreate() (fd int) {
	epfd, err := unix.EpollCreate1(0)
	if err != nil {
		fmt.Println(err)
	}
	return epfd
}

func (ep *Epoller) AddChannel(fd int,) {

	ev := unix.EpollEvent{}
	ev.Events = InEvents
	ev.Fd = int32(fd)

	err := unix.EpollCtl(ep.EpollFd, unix.EPOLL_CTL_ADD, fd, &ev)

	if err != nil {
		fmt.Println(err)
	}
	// ep.eventList = append(ep.eventList, ev)
	// ep.eventNum++

	ch := Channel{
		fd: fd,
	}
	ep.channelList = append(ep.channelList, ch)
}

func (ep *Epoller) RemoveChannel(fd int) {

}

func (ep *Epoller) Epoll(channels []Channel) int {

	evNum, err := unix.EpollWait(ep.EpollFd, ep.eventList, 10) // someone set the timeout to 0ms ，so the cpu conquer high to 15%，too tm cool

	if evNum > 0 {
		fmt.Printf("epoll here evnum %v", evNum)
	}
	if err != nil {
		fmt.Printf("epoll fd %v \n", ep.EpollFd)
		fmt.Println("here wrong")
		fmt.Println(err)
	}

	if evNum > 0 {

		ep.fillChannels(channels, evNum)

	}

	return evNum

}

func (ep *Epoller) fillChannels(channels []Channel, evNum int) {

	for i := 0; i < evNum; i++ {

		// append if there is no enough channels
		channels[i].setRevent(int(ep.eventList[i].Events))
		channels[i].fd = int(ep.eventList[i].Fd)

	}

}
