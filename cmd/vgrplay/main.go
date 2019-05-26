package main

import (
	"flag"
	"errors"
	vgr "github.com/rbxb/vaingloryreplay"
)

var (
	source string
	sname string
	overwrite string
	oname string
)

func init() {
	flag.StringVar(&source, "source", "./vainglory-replays", "The directory with the source vgr files. (./vainglory-replays)")
	flag.StringVar(&sname, "sname", "", "The name of the replay to play. (*picks the most recently modified replay*)")
	flag.StringVar(&overwrite, "overwrite", "./", "The directory with the active vgr files. (./)")
	flag.StringVar(&oname, "oname", "", "The name of the replay to overwrite. (*picks the most recently modified replay*)")
}

func main() {
	flag.Parse()
	if sname == "" {
		sname = vgr.LastModified(source)
		if sname == "" {
			panic(errors.New("No source replay available in " + source))
		}
	}
	if oname == "" {
		oname = vgr.LastModified(overwrite)
		if oname == "" {
			panic(errors.New("No overwrite replay available in " + overwrite))
		}
	}
	count := vgr.FragmentCount(overwrite, oname)
	for i := 0; i < count; i++ {
		if err := vgr.DeleteFragment(overwrite, oname, i); err != nil {
			panic(err)
		}
	}
	for i := 0; i < vgr.FragmentCount(source, sname); i++ {
		buf, err := vgr.ReadFragment(source, sname, i)
		if err != nil {
			panic(err)
		}
		if err := vgr.WriteFragment(overwrite, oname, i, buf); err != nil {
			panic(err)
		}
	}
}