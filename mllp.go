package mllp

import (
	"bufio"
	"io"

	"github.com/facebookgo/stackerr"
)

type ErrInvalidHeader error
type ErrInvalidTrailer error

func NewReader(r io.Reader) *Reader {
	return &Reader{
		b: bufio.NewReader(r),
	}
}

type Reader struct {
	b *bufio.Reader
}

func (r Reader) ReadMessage() ([]byte, error) {
	c, err := r.b.ReadByte()
	if err != nil {
		return nil, stackerr.Wrap(err)
	}

	if c != byte(0x0b) {
		return nil, ErrInvalidHeader(stackerr.Newf("invalid header found; expected 0x0b but got %02x", c))
	}

	d, err := r.b.ReadBytes(byte(0x1c))
	if err != nil {
		return nil, stackerr.Wrap(err)
	}

	t, err := r.b.ReadByte()
	if err != nil {
		return nil, stackerr.Wrap(err)
	}
	if t != byte(0x0d) {
		return nil, ErrInvalidTrailer(stackerr.Newf("invalid trailer found; expected 0x0d but got %02x", t))
	}

	return d[0 : len(d)-1], nil
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

type Writer struct {
	w io.Writer
}

func (w Writer) WriteMessage(b []byte) error {
	if _, err := w.w.Write([]byte{0x0b}); err != nil {
		return stackerr.Wrap(err)
	}

	for len(b) > 0 {
		n, err := w.w.Write(b)
		if err != nil {
			return stackerr.Wrap(err)
		}

		b = b[n:]
	}

	if _, err := w.w.Write([]byte{0x1c, 0x0d}); err != nil {
		return stackerr.Wrap(err)
	}

	return nil
}

func NewReadWriter(rw io.ReadWriter) *ReadWriter {
	return &ReadWriter{
		Reader: NewReader(rw),
		Writer: NewWriter(rw),
	}
}

type ReadWriter struct {
	*Reader
	*Writer
}
