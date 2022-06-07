package snet

import "goserve/internal/netbase"

type AcceptorLoop struct{

	ac *netbase.Acceptor
	acLoop *netbase.EventLoop
}


func NewAcLoop() (acl *AcceptorLoop){

	acl = new(AcceptorLoop)

	acl.ac = netbase.NewAcceptor()
    acl.acLoop = netbase.NewEventLoop()

	return
}

func defaultConnectionCallback(){
	



}

