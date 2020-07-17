package main

import (
	"github.com/rbxb/customprot"
	vgs "github.com/rbxb/vaingloryreplay/vainglorystream"
	vgr "github.com/rbxb/vaingloryreplay/vaingloryreplay"
	"fmt"
	"flag"
	"net/url"
	"strconv"
	"errors"
)

var (
	address string
	id string
	overwrite string
	oname string
)

func init() {
	flag.StringVar(&address, "address", "http://localhost:8080/view", "The address of the server. (http://localhost:8080/view)")
	flag.StringVar(&id, "id", "000000", "The stream id. (000000)")
	flag.StringVar(&overwrite, "overwrite", "./", "The directory with the active vgr files. (./)")
	flag.StringVar(&oname, "oname", "", "The name of the replay to overwrite. (*picks the most recently modified replay*)")
}

func main() {
	flag.Parse()
	if oname == "" {
		oname = vgr.LastModified(overwrite)
		if oname == "" {
			panic(errors.New("No overwrite replay available in " + overwrite))
		}
	}
	u, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Set("id", id)
	u.RawQuery = q.Encode()
	conn, err := customprot.Connect(u.String(), map[string]string {
		"vainglory-replay-stream": "view",
	}, map[string]string {
		"vainglory-replay-stream": "view",
	})
	if err != nil {
		panic(err)
	}

	for {
		frame, err := vgs.ReadFrame(conn)
		if err != nil || len(frame) == 0 {
			panic(err)
		}
		num, data, err := vgs.UnpackFrame(frame)
		if err != nil {
			panic(err)
		}
		seconds := (num + 1) * 10
		minutes := seconds / 60
		seconds = seconds % 60
		fmt.Println("Recieved frame " + strconv.Itoa(num) + " (through " + strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds) + ")")
		if err := vgr.WriteFrame(overwrite, oname, num, data); err != nil {
			panic(err)
		}
	}
}