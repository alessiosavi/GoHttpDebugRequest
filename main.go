package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
		fmt.Println(err)
	} else {
		fmt.Fprint(w, string(requestDump))
		fmt.Println(string(requestDump))
	}
}

func main() {
	port := flag.Int("port", 9999, "Bind the server to the given port")
	host := flag.String("host", "localhost", "Bind the server to the given host")
	flag.Parse()
	if *port == 9999 && *host == "localhost" {
		flag.Usage()
	}

	router := mux.NewRouter()
	router.HandleFunc("/get", DumpRequest).Methods("GET")
	router.HandleFunc("/post", DumpRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), router))
}
