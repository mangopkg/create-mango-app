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
}

func NewHandler(service *BookService) {

	h := BookHandler{
		MountAt:     "/book",
		BookService: *service,
	}

	f := reflect.TypeOf(&h)
	v := reflect.ValueOf(&h)

	h.BookService.SetupHandler(h.MountAt, f, v)
}

/*
<@route/{
"pattern": "/find",
"func": "Find",
"method": "GET"
}/route>
*/
func (h *BookHandler) Find() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Response.Message = "Successful"
		h.Response.StatusCode = 200
		h.Response.Data = h.BookService.GetBook()
		h.Response.Send(w)
	}
}
