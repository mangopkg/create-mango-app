package book

import "github.com/mangopkg/mango"

type BookService struct {
	*mango.Service
}

func NewService(s *mango.Service) {
	nS := BookService{
		s,
	}

	NewHandler(&nS)
}

func (s *BookService) GetBook() string {
	return "hello world book!"
}
