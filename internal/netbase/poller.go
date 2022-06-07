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
			fmt.Println(err)
		}

		for i := 0; i < event_num; i++ {

			ev := event_list[i]
			handler(ev.Fd, ev.Events)
		}
	}

}

func newPoller()(p *poller){
	return
}
