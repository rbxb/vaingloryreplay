package vainglorystream

import (
	"encoding/binary"
	"io"
	"errors"
	"bytes"
)

var ErrorClosed error = errors.New("Connection closed.")

type Frame []byte

func ReadFrame(r io.ReadCloser) (Frame, error) {
	var size uint32
	if err := binary.Read(r, binary.LittleEndian, &size); err != nil {
		r.Close()
		return nil, err
	}
	frame := make(Frame, int(size))
	read := 0
	for read < len(frame) {
		n, err := r.Read(frame[read:])
		if err != nil {
			r.Close()
			return nil, ErrorClosed
		}
		read += n
	}
	return frame, nil
}

func WriteFrame(w io.WriteCloser, frame Frame) error {
	if err := binary.Write(w, binary.LittleEndian, uint32(len(frame))); err != nil {
		w.Close()
		return err
	}
	if _, err := w.Write(frame); err != nil {
		w.Close()
		return err
	}
	return nil
}

func PackFrame(num int, data []byte) (Frame, error) {
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.LittleEndian, uint32(num)); err != nil {
		return nil, err
	}
	buf.Write(data)
	return buf.Bytes(), nil
}

func UnpackFrame(frame Frame) (int, []byte, error) {
	buf := bytes.NewBuffer(frame)
	var num uint32
	if err := binary.Read(buf, binary.LittleEndian, &num); err != nil {
		return 0, nil, err
	}
	return int(num), buf.Bytes(), nil
}