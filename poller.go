package snet

import (
	"fmt"
	"github.com/shinoshina/snet/internal/base/logger"
	"runtime"

	"golang.org/x/sys/unix"
)

type poller struct {
	epoller_fd int
}



func (p *poller) Epoll(handler epoll_callback)  {

	event_list := make([]unix.EpollEvent, 128)


	//when msec set -1, a empty epoll_wait will be blocked until a event come
	msec := -1
	for {
		//fmt.Println("blocked?")
		event_num, err := unix.EpollWait(p.epoller_fd, event_list, msec) // someone set the timeout to 0ms ，so the cpu conquer high to 15%，too tm cool
		if event_num ==0 || (event_num <0 && err == unix.EINTR ) {    // but when the msec is 0 , epoll_wait interrupted systemcall(unix.EINTR) will not occur
			msec = -1
			if event_num == -1 {
				fmt.Println(err)
			}
			runtime.Gosched()
			continue
		}

		msec = 0

		if event_num == -1 {

			fmt.Println("err")
		}

		for i := 0; i < event_num; i++ {
			ev := event_list[i]
			logger.Infof("events : %v",ev.Events)
			logger.Infof("%v",ev.Fd)
			handler(ev.Fd, ev.Events)
		}
	}

}

func(p *poller)register(fd int){

	ev := unix.EpollEvent{}
	ev.Events = InEvents | unix.EPOLLET
	ev.Fd = int32(fd)

	err := unix.EpollCtl(p.epoller_fd, unix.EPOLL_CTL_ADD, fd, &ev)

	if err != nil {
		fmt.Println(err)
	}

	logger.Infof("poll register fd : %v ",fd)

}

func(p *poller)unregister(fd int32){

	err := unix.EpollCtl(p.epoller_fd,unix.EPOLL_CTL_DEL,int(fd),nil)
	if err!= nil{
		fmt.Println(err)
	}

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

	logger.Infof("epoll fd created : %v",epfd)
	return epfd
}


