package kk

import "fmt"

type EventLoop struct {
	epoller        *Epoller
	activeChannels []Channel
}

func NewEventLoop(ep *Epoller) (evl *EventLoop) {

	evl = new(EventLoop)

	evl.epoller = ep
	evl.activeChannels = make([]Channel, 10)                    //FIX ME: here if events too many,consider appending
	fmt.Printf("eloop epoll fd %v\n",evl.epoller.EpollFd)
	return evl
}
func (evl *EventLoop) Loop() {

	for {

		evNum := evl.epoller.Epoll(evl.activeChannels)
		if evNum > 0 {
			fmt.Printf("evnum %v \n", evNum)
		}
		if evNum > 0 {
			for i := 0; i < evNum; i++ {
				fmt.Println("loop handle")
				evl.activeChannels[i].handleEvent()
			}
		}

	}
}
