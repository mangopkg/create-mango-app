package book

import (
	"net/http"
	"reflect"

	"github.com/mangopkg/mango"
)

type BookHandler struct {
	Response    mango.Response
	MountAt     string
	BookService BookService
	Methods     map[string]func(http.ResponseWriter, *http.Request)
}

func NewHandler(service BookService) {

	h := BookHandler{
		MountAt:     "/book",
		BookService: service,
		Methods:     make(map[string]func(http.ResponseWriter, *http.Request)),
	}

	f := reflect.TypeOf(&h)
	v := reflect.ValueOf(&h)

	h.BookService.SetupHandler(h.MountAt, f, v, h.Methods)
}

/*
<@route{
"pattern": "/find",
"func": "Find",
"method": "GET"
}>
*/
func (h *BookHandler) Find() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Response.Message = "Successful"
		h.Response.StatusCode = 200
		h.Response.Data = h.BookService.GetBook()
		h.Response.Send(w)
	}
}
