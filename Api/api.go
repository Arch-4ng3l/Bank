package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "github.com/Arch-4ng3l/Bank/Database"
	types "github.com/Arch-4ng3l/Bank/Types"
)

type Server struct {
	Addr  string
	Store database.Storage
}

func New(addr string, store database.Storage) *Server {
	return &Server{
		Addr:  addr,
		Store: store,
	}
}

func (s *Server) Run() error {

	http.HandleFunc("/api/login", apiFuncToHttpHandler(s.handleLogin))
	http.HandleFunc("/api/signup", apiFuncToHttpHandler(s.handleSignUp))
	http.HandleFunc("/api/transaction", apiFuncToHttpHandler(s.handleTransaction))

	return http.ListenAndServe(s.Addr, nil)
}

func (s *Server) handleSignUp(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	req := &types.SignUpRequest{}

	if err := decoder.Decode(req); err != nil {
		return nil
	}

	req.Password = createHash(req.Password)

	acc := s.Store.SignUp(req)

	if acc == nil {
		return fmt.Errorf("Invalid Sign Up Credentials")
	}

	return nil
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	req := &types.LoginRequest{}

	if err := decoder.Decode(req); err != nil {
		return err
	}

	encryptedPw := createHash(req.Password)

	acc := s.Store.Login(req, encryptedPw)

	if acc == nil {
		return fmt.Errorf("Invalid Login Informations")
	}

	return json.NewEncoder(w).Encode(acc)
}

func (s *Server) handleTransaction(w http.ResponseWriter, r *http.Request) error {
	req := &types.Transaction{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	s.Store.Transaction(req)

	return nil
}

type ApiFunction = func(w http.ResponseWriter, r *http.Request) error

func apiFuncToHttpHandler(f ApiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err)
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
		}
	}
}

func createHash(in string) string {
	hash := sha256.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}
