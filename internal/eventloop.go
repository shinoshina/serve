package internal

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


		evNum := evl.epoller.Epoll(evl.activeChannels)
		if evNum > 0 {
			for i := 0; i < evNum; i++ {
				evl.activeChannels[i].handleEvent()
			}
		}


	}
}
