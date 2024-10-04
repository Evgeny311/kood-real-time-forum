package server

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleGetReactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reactions, err := s.store.Reaction().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusFound, Response{
			Message: "Successful",
			Data:    reactions,
		})
	}
}

func (s *server) handleGetReactionsByParentID() http.HandlerFunc {
	type requestBody struct {
		ParentID string `json:"parent_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req requestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		reactions, err := s.store.Reaction().GetByParentID(req.ParentID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    reactions,
		})
	}
}

func (s *server) handleGetUserReactions() http.HandlerFunc {
	type requestBody struct {
		ParentID string `json:"parent_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(ctxUserID).(string)

		var req requestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		reactions, err := s.store.Reaction().GetByUserParentID(req.ParentID, userID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    reactions,
		})
	}
}

func (s *server) handleAddReactionsToParent() http.HandlerFunc {
	type requestBody struct {
		ParentID   string `json:"parent_id"`
		ReactionID string `json:"reaction_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(ctxUserID).(string)

		var req requestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err := s.store.Reaction().AddToParent(req.ParentID, req.ReactionID, userID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, Response{
			Message: "Successful",
			Data:    req.ReactionID,
		})
	}
}

func (s *server) handleRemoveReaction() http.HandlerFunc {
	type requestBody struct {
		ParentID   string `json:"parent_id"`
		ReactionID string `json:"reaction_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(ctxUserID).(string)

		var req requestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Reaction().RemoveFromParent(req.ParentID, req.ReactionID, userID); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    nil,
		})
	}
}

func (s *server) handleGetReactionsOptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reactions, err := s.store.Reaction().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusFound, reactions)
	}
}
