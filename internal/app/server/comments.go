package server

import (
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/router"
	"net/http"
)

func (s *server) handleCommentCreation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &models.Comment{}
		if err := json.NewDecoder(r.Body).Decode(c); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		userID := r.Context().Value(ctxUserID).(string)
		err := s.store.Comment().Create(c, userID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, Response{
			Message: "Successful",
			Data:    c,
		})
	}
}

func (s *server) handleRemoveComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := router.Param(r.Context(), "id")

		if err := s.store.Comment().Delete(id); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    nil,
		})
	}
}
