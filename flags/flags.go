package flags

import (
	"fmt"
	"os"
	"strconv"

	"github.com/felixrabe/limited-http-static-server/errors"
)

func usage() {
	fmt.Println("Usage: limited-http-static-server -p port directory")
	os.Exit(1)
}

func Parse() (port int, dir string) {
	if len(os.Args) != 4 {
		usage()
	}
	var err error
	if os.Args[1] != "-p" {
		usage()
	}
	if port, err = strconv.Atoi(os.Args[2]); err != nil {
		errors.Fatal(err)
	}
	dir = os.Args[3]
	return
}
