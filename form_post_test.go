package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func FormPost(writer http.ResponseWriter, request *http.Request){
	//lewat parsing buildin
	middleName:=request.PostFormValue("middlename")

	//bisa lewat parsing
	err:=request.ParseForm()
	if err != nil{
		panic(err)
	}


	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	fmt.Fprintf(writer,"Heloo %s middle %s, lastname %s",firstName,middleName,lastName)
}

func TestFormBody(t *testing.T) {
	requestBody := strings.NewReader("first_name=Vandy&middlename=ahmad&last_name=misry")
	request:=httptest.NewRequest(http.MethodPost,"http://localhost:8080",requestBody)
	request.Header.Add("Content-type","application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	FormPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
