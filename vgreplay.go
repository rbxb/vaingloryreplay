package vaingloryreplay

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"os"
)

func Path(path, name string, frag int) string {
	return filepath.Join(path, name + "." + strconv.Itoa(frag) + ".vgr")
}
//returns the filepath to the .vgr file

func Name(str string) string {
	ext := filepath.Ext(str)
	if ext == ".vgr" {
		str = str[:len(str)-len(ext)]
		num := filepath.Ext(str)
		str = str[:len(str)-len(num)]
		return str
	} else {
		return ""
	}
}
//returns the name of the replay from the full file name

func ListReplays(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}
	}
	names := make([]string, 0)
	for _, info := range files {
		name := Name(info.Name())
		if name == "" {
			continue
		}
		exists := false
		for _, n := range names {
			if n == name {
				exists = true
				break
			}
		}
		if !exists {
			names = append(names, name)
		}
	}
	return names
}
//returns a list of the names of every replay in the directory

func LastModified(path string) string {
	name := ""
	if names := ListReplays(path); len(names) > 0 {
		var max int64 = -1
		for _, n := range names {
			info, err := os.Stat(Path(path, n, 0))
			if err != nil {
				panic(err)
			}
			t := info.ModTime().UnixNano()
			if t > max {
				max = t
				name = n
			}
		}
	}
	return name
}
//returns the name of the most recently modified replay

func FragmentCount(path string, name string) int {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0
	}
	count := 0
	for _, info := range files {
		if Name(info.Name()) == name {
			count++
		}
	}
	return count
}
//returns the number of fragments the replay has

func ReadFragment(path, name string, frag int) (* bytes.Buffer, error) {
	b, err := ioutil.ReadFile(Path(path,name,frag))
	return bytes.NewBuffer(b), err
}
//reads a fragment

func WriteFragment(path, name string, frag int, buf * bytes.Buffer) error {
	path = Path(path,name,frag)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	if err = f.Truncate(int64(buf.Len())); err != nil {
		return err
	}
	if _, err = buf.WriteTo(f); err != nil {
		return err
	}
	return nil
}
//writes or overwrites a fragment

func DeleteFragment(path, name string, frag int) error {
	return os.Remove(Path(path,name,frag))
}
//deletes a fragment