package netbase


type connction struct{

	evl_from *eventloop
	buf string // use string temply
	fd int32
	raddr raw_address
}

func newConnection(fd int32) (conn *connction){

	conn = new(connction)
    conn.fd = fd
	// conn.evl_from = evl

	return;
}
