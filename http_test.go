package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello",nil)
	recorder := httptest.NewRecorder()
	HelloHandler(recorder,request)
	response:=recorder.Result()
	body, _:=io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
