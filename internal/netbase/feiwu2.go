package netbase

import "fmt"

type EventLoop struct {
	epoller        *Epoller
	activeChannels []Channel
}

func NewEventLoop() (evl *EventLoop) {

	evl = new(EventLoop)
	evl.activeChannels = make([]Channel, 10)                    //FIX ME: here if events too many,consider appending
	evl.epoller = NewEpoller()


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


func(evl *EventLoop) Register(fd int,channelType int){

	evl.epoller.AddChannel(fd)

}
