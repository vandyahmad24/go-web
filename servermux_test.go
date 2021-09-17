package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Hello World")
	})

	mux.HandleFunc("/vandy/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Vandy")
	})

	mux.HandleFunc("/vandy/ahmad/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Vandy Ahmad")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err:=server.ListenAndServe()
	if err!=nil{panic(err)}

}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,request.Method)
		fmt.Fprintln(writer,request.RequestURI)
	}


	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}
	err:=server.ListenAndServe()
	if err!=nil{panic(err)}
}
