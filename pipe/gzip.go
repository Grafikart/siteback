package siteback

import (
	"compress/gzip"
	"io"
)

func Gzip(input io.Reader, filename string) (io.Reader, error) {
	r, w := io.Pipe()
	zip, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
	zip.Name = filename
	go func() {
		io.Copy(zip, input)
		defer w.Close()
		defer zip.Close()
	}()
	return r, nil
}
