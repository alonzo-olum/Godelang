package main

var (
	freeList = make(chan *Buffer, 100)
	serverChan = make(chan *Buffer)
)

func client() {
	for {
		var b *Buffer
		//grab a buffer if available; allocate if not
		select {
		case b = <-freeList:
			// got one; to-do
		default:
			b = new(Buffer)
		}
		load(b)	//read next message from the buffer
		serverChan<- b	//send to server
	}
}

func server() {
	for {
		b := <-serverChan
		process(b)
		//re-use buffer if there's room
		select {
		case freeList<- b:
			//buffer on freeList; to-do
		default:
			//freeList is full; thrown to gb
		}
	}
}
