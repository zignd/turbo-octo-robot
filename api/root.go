package api

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hi there, welcome to /")
}
