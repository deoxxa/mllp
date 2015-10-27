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

	d, err := r.b.ReadBytes(byte(0x1C))
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
