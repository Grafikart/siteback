package siteback

import (
	"io"
	"os"
	"os/exec"
	"regexp"
)

// Génère un dump SQL à partir du chemin fournit par symfony
func DumpDB() (r io.Reader, err error) {
	regex, _ := regexp.Compile(`mysql://([^:]+):([^@]+)[^/]+/(\w*)`)
	matches := regex.FindStringSubmatch(os.Getenv("DATABASE_URL"))
	cmd := exec.Command("mysqldump", "-u"+matches[1], "-p"+matches[2], matches[3])
	r, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	return r, nil
}
