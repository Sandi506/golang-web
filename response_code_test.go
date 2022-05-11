package golang_web

import (
	"fmt"
	"io"
	hhttp "net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer hhttp.ResponseWriter, request *hhttp.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(hhttp.StatusBadRequest)
		fmt.Fprint(writer, " name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Sandi", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
