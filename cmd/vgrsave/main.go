package main

import (
	"flag"
	"os"
	"errors"
	"strconv"
	"time"
	vgr "github.com/rbxb/vaingloryreplay"
)

var (
	source string
	name string
	save string
	sname string
)

func init() {
	flag.StringVar(&source, "source", "./", "The directory with the source vgr files. (./)")
	flag.StringVar(&name, "name", "", "The name of the replay to save. (*picks the most recently modified replay*)")
	flag.StringVar(&save, "save", "./vainglory-replays", "The directory where the replay will be saved. (./vainglory-replays)")
	flag.StringVar(&sname, "sname", "", "The name to save the replay as. (*auto generated*)")
}

func main() {
	flag.Parse()
	if name == "" {
		name = vgr.LastModified(source)
		if name == "" {
			panic(errors.New("No source replay available in " + source))
		}
	}
	if sname == "" {
		sname = "replay-" + strconv.Itoa(int(time.Now().UnixNano()))
	}
	if err := os.MkdirAll(save, 0755); err != nil {
		panic(err)
	}
	for i := 0; i < vgr.FrameCount(source, name); i++ {
		b, err := vgr.ReadFrame(source, name, i)
		if err != nil {
			panic(err)
		}
		if err := vgr.WriteFrame(save, sname, i, b); err != nil {
			panic(err)
		}
	}
}