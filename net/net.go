package net

import (
	"net/http"
	"strconv"

	"github.com/felixrabe/limited-http-static-server/errors"
)

func Serve(port int, dir string) {
	addr := ":" + strconv.Itoa(port)
	handler := http.FileServer(http.Dir(dir))
	if err := http.ListenAndServe(addr, handler); err != nil {
		errors.Fatal(err)
	}
}
