package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHanlder(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic nya
		fmt.Fprint(writer, "Hello Sandi")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Sandi")
	})
	mux.HandleFunc("/hai", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hai")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images")
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)

	}
}

func TestRequest(t *testing.T) { // Get / isi/isi/
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
