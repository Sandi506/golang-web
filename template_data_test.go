package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Sandi",
		"Address": map[string]interface{}{
			"Street": "Jalan Mara Bahaya",
		},
	})

}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	reocrder := httptest.NewRecorder()

	TemplateDataMap(reocrder, request)

	body, _ := io.ReadAll(reocrder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}
type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruck(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struck",
		Name:  "Sandi",
		Address: Address{
			Street: "Jalan Veteran",
		},
	})
}

func TestTemplateDataStruck(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruck(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
