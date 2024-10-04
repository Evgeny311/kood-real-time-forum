package server

import "net/http"

func (s *server) handleGetAllCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := s.store.Category().GetAll()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    categories,
		})
	}
}
