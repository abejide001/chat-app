//main.go
package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(w, r)
	})
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
