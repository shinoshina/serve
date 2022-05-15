package internal

import (
	"fmt"

	"golang.org/x/sys/unix"
)

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

const (
	ErrEvents = unix.EPOLLERR | unix.EPOLLHUP | unix.EPOLLRDHUP

	OutEvents = ErrEvents | unix.EPOLLOUT

	InEvents = ErrEvents | unix.EPOLLIN | unix.EPOLLPRI
)

type Epoller struct {
	thisLoop    *EventLoop
	channelList []Channel
	sequence    uint16
	epollFd     int
	eventList   []unix.EpollEvent
	eventNum    int
}

func EpollCreate() (fd int) {
	epfd, err := unix.EpollCreate1(0)
	if err != nil {
		fmt.Println(err)
	}
	return epfd
}

func (ep Epoller) AddChannel(fd int) {

	ev := unix.EpollEvent{}

	unix.EpollCtl(ep.epollFd, unix.EPOLL_CTL_ADD, fd, &ev)

	// ep.eventList = append(ep.eventList, ev)
	// ep.eventNum++

	ch := Channel{
		fd: fd,
	}
	ep.channelList = append(ep.channelList, ch)
}

func (ep Epoller) RemoveChannel(fd int) {

}

func (ep Epoller) Epoll(channels []Channel) int {

	evNum, err := unix.EpollWait(ep.epollFd, ep.eventList, 10)

	if err != nil {
		fmt.Println(err)
	}

	if evNum > 0 {

		ep.fillChannels(channels,evNum)

	}

	return evNum

}

func (ep Epoller) fillChannels(channels []Channel, evNum int) {

	for i := 0; i < evNum; i++ {

		tmpch := Channel{
		}
		tmpch.setRevent(int(ep.eventList[i].Events))
		tmpch.fd = int(ep.eventList[i].Fd)

		channels = append(channels, tmpch)

	}

}
