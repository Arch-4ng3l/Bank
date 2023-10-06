package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	types "github.com/Arch-4ng3l/Bank/Types"
)

type Server struct {
	ListendingAddr string
}

func New(addr string) *Server {
	return &Server{
		addr,
	}
}

func (s *Server) Run() error {

	http.HandleFunc("/api/login", apiFuncToHttpHandler(s.handleLogin))
	http.HandleFunc("/api/signup", apiFuncToHttpHandler(s.handleSignUp))

	return http.ListenAndServe(s.ListendingAddr, nil)
}

func (s *Server) handleSignUp(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	req := &types.SignUpRequest{}
	if err := decoder.Decode(req); err != nil {
		return nil
	}

	req.Print()

	account := &types.Account{
		Username:     req.Username,
		EmailAddress: req.EmailAddress,
		Password:     req.Password,
		CreatedAt:    time.Now(),
	}

	account.Print()

	return nil
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {

	decoder := json.NewDecoder(r.Body)
	req := &types.LoginRequest{}

	if err := decoder.Decode(req); err != nil {
		return err
	}

	return nil
}

type ApiFunction = func(w http.ResponseWriter, r *http.Request) error

func apiFuncToHttpHandler(f ApiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err)
		}
	}
}
