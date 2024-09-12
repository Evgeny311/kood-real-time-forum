package server

import (
	"encoding/json"
	"net/http"
	"time"

	"forum/internal/models"
	"forum/pkg/router"
)

func (s *server) handleCreateChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		otherUserID := router.Param(r.Context(), "user_id")

		userExists, err := s.store.User().IsUser(otherUserID)
		if err != nil {
			s.respond(w, r, http.StatusBadRequest, err)
			return
		} else if !userExists {
			s.respond(w, r, http.StatusBadRequest, Response{
				Message: "The user by this ID doesn't exist",
			})
			return
		}

		userID := r.Context().Value(ctxUserID).(string)

		if userID == otherUserID {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		res, err := s.store.Chat().CheckChatExists(userID, otherUserID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var chat_id []string
		if res == nil {
			chat, err := s.store.Chat().Create(userID, otherUserID)
			if err != nil {
				s.error(w, r, http.StatusCreated, err)
				return
			}
			chat_id = append(chat_id, chat.ID)
			s.respond(w, r, http.StatusOK, Response{
				Message: "Successfully created chat",
				Data: map[string][]string{
					"chat_id": chat_id,
				},
			})
			return
		} else {
			chat_id = res
		}

		// fmt.Println(res)
		s.respond(w, r, http.StatusOK, Response{
			Message: "Successfully found chat",
			Data: map[string][]string{
				"chat_id": chat_id,
			},
		})
		return
	}
}

func (s *server) handleWriteLines() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var line models.Line
		if err := json.NewDecoder(r.Body).Decode(&line); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		line.UserID = r.Context().Value(ctxUserID).(string)

		if err := s.store.Chat().WriteLine(line); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, Response{
			Message: "Line Written",
		})
	}
}

func (s *server) handleInitChatLines() http.HandlerFunc {
	type requestBody struct {
		Chat_id string    `json:"chat_id"`
		Time    time.Time `json:"timestamp"`
		Count   int       `json:"count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		limit := 20
		var req requestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		offset := req.Count * limit
		lines, err := s.store.Chat().GetLinesInit(req.Chat_id, req.Time, limit, offset)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, Response{
			Message: "Lines gotten successfully",
			Data:    lines,
		})
	}
}
