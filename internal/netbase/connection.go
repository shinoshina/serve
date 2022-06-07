package netbase


type connction struct{

	evl_from *eventloop
	buf string // use string temply
	fd int32
}

func newConnection(fd int32,evl *eventloop) (conn *connction){

	conn = new(connction)
    conn.fd = fd
	conn.evl_from = evl

	return;
}
