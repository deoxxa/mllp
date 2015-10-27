package mllp

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func wrapWithMarkers(b []byte) []byte {
	return append(append([]byte{0x0b}, b...), 0x0d, 0x1c, 0x0d)
}

func TestReadMessage(t *testing.T) {
	r := NewReader(bytes.NewReader(wrapWithMarkers([]byte("hello"))))

	m, err := r.ReadMessage()
	if err != nil {
		panic(err)
	}

	if string(m) != "hello" {
		panic(fmt.Errorf("data was corrupted"))
	}
}

func TestWriter(t *testing.T) {
	b := bytes.NewBuffer(nil)

	w := NewWriter(b)

	if err := w.WriteMessage([]byte("hello")); err != nil {
		panic(err)
	}

	if !bytes.Equal(b.Bytes(), wrapWithMarkers([]byte("hello"))) {
		panic(fmt.Errorf("data was corrupted"))
	}
}

func TestReadWriter(t *testing.T) {
	r := bytes.NewReader(wrapWithMarkers([]byte("input")))
	w := bytes.NewBuffer(nil)

	brw := bufio.NewReadWriter(bufio.NewReader(r), bufio.NewWriter(w))

	rw := NewReadWriter(brw)

	m, err := rw.ReadMessage()
	if err != nil {
		panic(err)
	}

	if string(m) != "input" {
		panic(fmt.Errorf("data was corrupted"))
	}

	if err := rw.WriteMessage([]byte("output")); err != nil {
		panic(err)
	}

	brw.Flush()

	if !bytes.Equal(w.Bytes(), wrapWithMarkers([]byte("output"))) {
		panic(fmt.Errorf("data was corrupted"))
	}
}
