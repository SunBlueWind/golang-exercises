package limitreader

import "io"

type reader struct {
	r     io.Reader
	cur   int64
	limit int64
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit-r.cur])
	r.cur += int64(n)
	if r.cur >= r.limit {
		err = io.EOF
	}
	return
}

// LimitReader returns a Reader that reads from r,
// but reports an EOF condition after n bytes
func LimitReader(r io.Reader, n int64) io.Reader {
	return &reader{r: r, limit: n}
}
