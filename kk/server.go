package kk



type Option struct{

}
type Server struct{


	option Option
	accptor Acceptor
	mainLoop EventLoop
	subLoops []EventLoop
	connMap  map[TcpAddress]int    //  address : loop number
	profile TcpAddress



}