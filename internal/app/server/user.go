package server

import (
	"encoding/json"
	"errors"
	"forum/internal/models"
	"forum/pkg/jwttoken"
	"forum/pkg/router"
	"net/http"
	"os"
	"time"
)

func (s *server) handleLogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deletedCookie := s.deleteCookie()
		http.SetCookie(w, &deletedCookie)
		s.respond(w, r, http.StatusOK, Response{Message: "Successful logged out"})
	}
}

func (s *server) handleUsersLogin() http.HandlerFunc {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := s.store.User().Check(requestBody.Email)
		if err != nil || !user.ComparePassword(requestBody.Password) {
			s.error(w, r, http.StatusUnauthorized, errors.New("invalid login credentials"))
			return
		}

		expiration := time.Now().Add(5 * time.Hour)
		alg := jwttoken.HmacSha256(os.Getenv(jwtKey))
		claims := jwttoken.NewClaims(user.ID, expiration.Unix())
		token, err := alg.Encode(claims)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		cookie := http.Cookie{
			Name:     sessionName,
			Value:    token,
			Expires:  expiration,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		}

		http.SetCookie(w, &cookie)

		user.Sanitize()
		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    user,
		})
	}
}

func (s *server) handleUsersGetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from cookie
		id := router.Param(r.Context(), "id")

		if id == "" {
			id = r.Context().Value(ctxUserID).(string)
		}
		user, err := s.store.User().FindByID(id)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    user,
		})
	}
}

func (s *server) handleUsersDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := router.Param(r.Context(), "id")
		err := s.store.User().Delete(id)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    nil,
		})
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.User().Create(user); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, Response{
			Message: "Successful",
			Data:    user,
		})
	}
}

func (s *server) handleUsersGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(ctxUserID).(string)

		data, err := s.store.User().GetAllOtherUsers(userID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, Response{
			Message: "Successful",
			Data:    data,
		})
	}
}
