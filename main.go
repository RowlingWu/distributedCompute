package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	server := newServer()
	server.Run(":8080")
}

func newServer() *negroni.Negroni {
	mx := mux.NewRouter()
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	fmt.Println("path:", webRoot)

	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/data/"))))

	n := negroni.Classic()
	n.UseHandler(mx)
	return n
}
