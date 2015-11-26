package main

import (
	"io"
	"os"
	"regexp"

	"fknsrs.biz/p/mllp"
	"github.com/facebookgo/stackerr"
)

func main() {
	rx := regexp.MustCompile(os.Args[1])

	r := mllp.NewReader(os.Stdin)
	w := mllp.NewWriter(os.Stdout)

	i := 0
	for {
		m, err := r.ReadMessage()
		if err != nil {
			if stackerr.HasUnderlying(err, stackerr.Equals(io.EOF)) {
				break
			}

			panic(err)
		}

		if rx.Match(m) {
			if err := w.WriteMessage(m); err != nil {
				panic(err)
			}
		}

		i++
	}
}
