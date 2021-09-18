package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)
//go:embed resource/ok.html
var resourceOk string

//go:embed resource/notfound.html
var resourceNotFound string

func ServeFile(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		fmt.Fprint(w,resourceOk)
	}else{
		fmt.Fprint(w,resourceNotFound)
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:9191",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}