package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `
<html><body><form action="#" method="post" name="bar">
	<input type="text" name="in"/>
	<input type="text" name="int"/>
	<input type="submit" value="Submit"/>
</form></body></html>
`

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, world.</h1>")
	panic("test1 panic")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v\n", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("listen and server failed, err: ", err)
	}
}
