package main

import (
	"github.com/rbxb/customprot"
	vgr "github.com/rbxb/vaingloryreplay/vaingloryreplay"
	vgs "github.com/rbxb/vaingloryreplay/vainglorystream"
	"fmt"
	"time"
	"flag"
	"errors"
	"net"
	"strconv"
)

var (
	source string
	name string
	address string
)

func init() {
	flag.StringVar(&source, "source", "./", "The directory with the source vgr files. (./)")
	flag.StringVar(&name, "name", "", "The name of the replay to save. (*picks the most recently modified replay*)")
	flag.StringVar(&address, "address", "http://localhost:8080/stream", "The address of the server. (http://localhost:8080/stream)")
}

func main() {
	flag.Parse()
	if name == "" {
		name = vgr.LastModified(source)
		if name == "" {
			panic(errors.New("No source replay available in " + source))
		}
	}

	conn, err := customprot.Connect(address, map[string]string {
		"vainglory-replay-stream": "stream",
	}, map[string]string {
		"vainglory-replay-stream": "stream",
	})
	if err != nil {
		panic(err)
	}

	b, err := vgs.ReadFrame(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("----- STREAM CODE -----\r\n" + string(b) + "\r\n-----------------------")

	go func(conn net.Conn){
		for {
			b, err := vgs.ReadFrame(conn)
			if err != nil {
				panic(err)
			}
			fmt.Println("Recieved: " + string(b))
		}
	}(conn)

	count := 0
	sent := 0
	for {
		count = vgr.FrameCount(source, name)
		for ; sent < count - 1; sent++ {
			data, err := vgr.ReadFrame(source, name, sent)
			if err != nil {
				panic(err)
			}
			frame, err := vgs.PackFrame(sent, data)
			if err != nil {
				panic(err)
			}
			if err := vgs.WriteFrame(conn, frame); err != nil {
				panic(err)
			}
			fmt.Println("Sent frame " + strconv.Itoa(sent) + " (" + strconv.Itoa(sent+1) + " of " + strconv.Itoa(count) + ")")
		}
		time.Sleep(500 * time.Millisecond)
	}
}