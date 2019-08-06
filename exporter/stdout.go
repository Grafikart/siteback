package siteback

import (
	"fmt"
	"io"
)

func Stdout(r io.Reader) (string, error) {
	p := make([]byte, 4)

	for {
		n, err := r.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remainding bytes.
				break
			}
			return "", err
		}
	}
	return "ok", nil
}
