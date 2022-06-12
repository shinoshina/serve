package netbase


type connction struct{

	evl_from *eventloop
	buf []byte // use string temply
	fd int32
	local_addr rawAddress
	remote_addr rawAddress
 }

func newConnection(fd int32) (c *connction){

	c = new(connction)
    c.fd = fd
	c.buf = make([]byte, 50)

	return;
}

func (c *connction)write(){


}
func (c *connction)read(){
	
}
