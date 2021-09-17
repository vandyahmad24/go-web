package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request)  {
	name:=request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(writer,"Helo")
	}else{
		fmt.Fprintf(writer, "Hello %s",name)
	}
}
func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request){
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s",firstName,lastName)
}

func MultipleValueQueryParamter(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, "_"))
}


func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/vandy?name=vandy&name=ahmad&name=misry",nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParamter(recorder,request)

	response := recorder.Result()
	body, _ :=io.ReadAll(response.Body)
	fmt.Println(string(body))
}