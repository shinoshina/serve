package internal




type EventCallback func ()(state string)


var readCb EventCallback
var	writeCb EventCallback
var	errCb EventCallback

type Channel struct{
	fd int
	revent int
}


func (ch Channel) setRevent(event int){
	ch.revent = event
}

func (ch Channel) handleEvent(){

	if ch.revent & InEvents != 0{

			readCb()
		
	}

	// }else if ch.revent & OutEvents != 0{

	// 	if ch.writeCb != nil{
	// 		ch.writeCb()
	// 	}
	// }
}

