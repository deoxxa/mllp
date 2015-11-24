package main

import (
	"fmt"
	"io"
	"os"

	"fknsrs.biz/p/mllp"
	"github.com/facebookgo/stackerr"
)

func main() {
	r := mllp.NewReader(os.Stdin)

	i := 0
	for {
		if _, err := r.ReadMessage(); err != nil {
			if stackerr.HasUnderlying(err, stackerr.Equals(io.EOF)) {
				break
			}

			panic(err)
		}

		i++
	}

	fmt.Printf("total messages: %d\n", i)
}
