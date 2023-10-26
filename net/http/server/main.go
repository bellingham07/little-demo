package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("我收到了，给你返回get"))
		break
	case "POST":
		b, _ := io.ReadAll(r.Body)
		header := w.Header()
		header["test"] = []string{"test1", "test2"}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		break

	default:
		w.Write([]byte("我收到了，给你返回,都不是"))
		break

	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler two"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	mux.HandleFunc("/test2", handler2)
	http.ListenAndServe(":8080", mux)
}
