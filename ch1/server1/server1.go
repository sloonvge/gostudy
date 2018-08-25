package main

import (
	"net/http"
	"log"
	"fmt"
)

/**
* Created by wanjx on 2018/8/25.
*/

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}




