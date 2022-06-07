package netbase

import (
	"fmt"

	"golang.org/x/sys/unix"
)

type poller struct {
	epoller_fd int
}

func (p *poller) Epoll(handler epoll_callback)  {

	event_list := make([]unix.EpollEvent, 128)

	for {

		event_num, err := unix.EpollWait(p.epoller_fd, event_list, 10) // someone set the timeout to 0ms ，so the cpu conquer high to 15%，too tm cool
		if err != nil {
			fmt.Println("epoll wait error")
			fmt.Println(err)
		}

		if event_num == -1 {

			fmt.Println("err")
		}

		for i := 0; i < event_num; i++ {

			ev := event_list[i]
			fmt.Println("is it true?")
			fmt.Println(ev.Events)
			fmt.Println(ev.Fd)
			handler(ev.Fd, ev.Events)
		}
	}

}

func(p *poller)register(fd int){

	ev := unix.EpollEvent{}
	ev.Events = InEvents
	ev.Fd = int32(fd)

	err := unix.EpollCtl(p.epoller_fd, unix.EPOLL_CTL_ADD, fd, &ev)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("poll register fd : %v \n",fd)

}

func newPoller()(p *poller){

	p = new(poller)
	p.epoller_fd = epoll_create()

	return
}

func epoll_create() (fd int) {
	epfd, err := unix.EpollCreate1(0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("epoll fd created : %v\n",epfd)
	return epfd
}


