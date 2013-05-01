package debugreader

import (
	"fmt"
	"io"
	"os"
)

type debugReader struct {
	r io.Reader
}

func NewReader(r io.Reader) io.Reader {
	return &debugReader{r: r}
}

func (r *debugReader) Read(buf []byte) (n int, err error) {
	fmt.Fprintf(os.Stderr, "buf in: %v\n", buf)
	n, err = r.r.Read(buf)
	fmt.Fprintf(os.Stderr, "n: %v\n", n)
	fmt.Fprintf(os.Stderr, "err: %v\n", err)
	fmt.Fprintf(os.Stderr, "buf out: %v\n", buf)
	return
}
