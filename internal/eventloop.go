package internal

import "fmt"

type EventLoop struct {
	epoller        *Epoller
	activeChannels []Channel
}


func NewEventLoop(ep *Epoller)(evl *EventLoop){


	evl = new(EventLoop)

	evl.epoller = ep
	evl.activeChannels = make([]Channel, 0)
	return evl
}
func (evl EventLoop) Loop() {

	for {


		fmt.Println("loop begin")
		evNum := evl.epoller.Epoll(evl.activeChannels)
		fmt.Printf("evnum %v \n",evNum)
		if evNum > 0 {
			for i := 0; i < evNum; i++ {
				fmt.Println("loop handle")
				evl.activeChannels[i].handleEvent()
			}
		}
		fmt.Println("loop end")


	}
}
