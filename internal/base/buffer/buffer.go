package buffer


type Buffer struct{

	Raw []byte
	S int
	E int
}

func NewBuffer(size int)(b *Buffer){

	b = new(Buffer)
	b.Raw = make([]byte, size)
	b.E = 0
	b.S = 0
	return
}