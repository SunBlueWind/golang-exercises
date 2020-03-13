package countingwriter

import "io"

type byteCounter struct {
	w     io.Writer
	bytes int64
}

func (b *byteCounter) Write(p []byte) (int, error) {
	n, err := b.Write(p)
	if err != nil {
		return 0, err
	}
	b.bytes += int64(n)
	return n, nil
}

// CountingWriter returns a new io.Writer that wraps the original,
// and a pointer that contains the number of bytes written to the
// new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := byteCounter{w, 0}
	return &counter, &counter.bytes
}
