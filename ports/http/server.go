package http

import (
	"app/app/model"
	"app/app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	Port     string
	useCases *usecase.UseCase
}

func NewServer(port string, useCases *usecase.UseCase) *Server {
	return &Server{
		Port:     port,
		useCases: useCases,
	}
}

func (s *Server) Run() {
	http.HandleFunc("/sign-in", handleSignIn(s.useCases))

	log.Fatal(http.ListenAndServe(s.Port, nil))
}

func handleSignIn(useCase *usecase.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials model.Credentials

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		session, err := useCase.SignIn(credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   session.Token,
			Expires: session.ExpiresAt,
		})
	}
}
