package web

import "net/http"

func AddFrontend() {
	fs := http.FileServer(http.Dir("./Web/src"))
	http.Handle("/", fs)
}
