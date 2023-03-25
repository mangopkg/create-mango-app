package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mangopkg/create-mango-app/book"
	"github.com/mangopkg/mango"
)

func Launch() {
	r := chi.NewRouter()

	s := *mango.New(mango.ServiceInit{R: r})

	s.R.Use(middleware.Logger)

	initServices(&s)

	s.GenerateDocs()

	log.Println("Serving on port 3000")

	if err := http.ListenAndServe("localhost:3000", r); err != nil {
		log.Println("ERROR LAUNCHING API: ", err)
	}

}

func initServices(s *mango.Service) {
	book.NewService(s)
}
