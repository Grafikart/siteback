package siteback

import (
	"io"
	"io/ioutil"
)

func File(r io.Reader, filename string) (string, error) {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filename, s, 0644)
	if err != nil {
		return "", err
	}
	return "ok", nil
}
