package streamcore

import (
	"io"
	"fmt"
	vgs "github.com/rbxb/vaingloryreplay/vainglorystream"
)

type Child struct {
	w io.WriteCloser
	lock chan byte
}

func NewChild(w io.WriteCloser) * Child {
	return &Child{
		w: w,
		lock: make(chan byte, 1),
	}
}

func(child * Child) Send(frame vgs.Frame) bool {
	child.lock <- 0
	num, _, _ := vgs.UnpackFrame(frame)
	fmt.Println(num)
	err := vgs.WriteFrame(child.w, frame)
	if err == nil {
		fmt.Println("CHILD:  Sent frame.")
	} else {
		fmt.Println("CHILD:  Send failed.")
	}
	<- child.lock
	return err == nil
}