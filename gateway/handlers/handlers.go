package handlers

import "net/http"

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (s *Repository) SignUp(w http.ResponseWriter, r *http.Request) {

}
