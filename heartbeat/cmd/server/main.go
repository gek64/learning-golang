package main

// http://www.suoniao.com/topic/601257e2e1dc400ae800c2c9

var (
	Req_REGISTER byte = 1
	Res_REGISTER byte = 2

	Req_HEARTBEAT byte = 3
	Res_HEARTBEAT byte = 4

	Req byte = 5
	Res byte = 6

	CMap map[string]*CS
)

type CS struct {
	Rch chan []byte
	Wch chan []byte
	Dch chan []byte
	Uid string
}

func NewCs(uid string) *CS {
	return &CS{
		Rch: make(chan []byte),
		Wch: make(chan []byte),
		Dch: make(chan []byte),
		Uid: uid,
	}
}

func main() {
}
