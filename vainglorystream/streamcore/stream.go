package streamcore

import (
	"fmt"
	"io"
	vgs "github.com/rbxb/vaingloryreplay/vainglorystream"
)

type CloseFunc func(* Stream)

type Stream struct {
	r io.ReadCloser
	lock chan byte
	children [] * Child
	frames []vgs.Frame
	close CloseFunc
}

func NewStream(r io.ReadCloser, close CloseFunc) * Stream {
	stream := &Stream{
		r: r,
		lock: make(chan byte, 1),
		children: make([] * Child, 0),
		frames: make([]vgs.Frame, 0),
		close: close,
	}
	go func(stream * Stream){
		fmt.Println("STREAM: New stream listening.")
		for {
			fmt.Println("STREAM: Waiting for frame.")
			if frame, err := vgs.ReadFrame(stream.r); err == nil {
				fmt.Println("STREAM: Frame recieved.")
				stream.lock <- 0
				if stream.children != nil {
					stream.frames = append(stream.frames, frame)
					for i := 0; i < len(stream.children); i++ {
						fmt.Println("STREAM: Sending to child.")
						if !stream.children[i].Send(frame) {
							stream.children = append(stream.children[:i], stream.children[i+1:]...)
							i--
							fmt.Println("STREAM: Send failed; removed child from stream.")
						}
					}
				}
				<- stream.lock
			} else {
				fmt.Println("STREAM: Recieve failed.")
				stream.Close()
				break
			}
		}
	}(stream)
	return stream
}

func(stream * Stream) Join(child * Child) {
	stream.lock <- 0
	if stream.children != nil {
		stream.children = append(stream.children, child)
		fmt.Println("STREAM: Added child to stream.")
	}
	frames := make([]vgs.Frame,len(stream.frames))
	copy(frames, stream.frames)
	go func(child * Child, frames []vgs.Frame){
		fmt.Println("STREAM: Bringing new child up to date.")
		for _, frame := range frames {
			fmt.Println("STREAM: Sending child old frame.")
			if !child.Send(frame) {
				fmt.Println("STREAM: Send failed.")
				break
			}
		}
	}(child,frames)
	<- stream.lock
}

func(stream * Stream) Ok() bool {
	stream.lock <- 0
	ok := stream.children != nil
	<- stream.lock
	return ok
}

func(stream * Stream) Close() {
	stream.lock <- 0
	if stream.children != nil {
		stream.r.Close()
		for _, child := range stream.children {
			child.w.Close()
		fmt.Println("STREAM: Closed child.")
		}
		stream.children = nil
	}
	stream.close(stream)
	fmt.Println("STREAM: Closed stream.")
	<- stream.lock
}