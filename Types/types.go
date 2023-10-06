package types

import (
	"fmt"
	"time"
)

type SignUpRequest struct {
	Username     string `json:"username"`
	EmailAddress string `json:"email"`
	Password     string `json:"password"`
}

func (sr *SignUpRequest) Print() {
	fmt.Println("Username: " + sr.Username)
	fmt.Println("Email: " + sr.EmailAddress)
	fmt.Println("Password: " + sr.Password)
}

type Account struct {
	Username     string    `json:"username"`
	EmailAddress string    `json:"email"`
	Password     string    `json:""`
	CreatedAt    time.Time `json:"createdAt"`
	Value        int
}

func (acc *Account) Print() {
	fmt.Println("Username: " + acc.Username)
	fmt.Println("Email: " + acc.EmailAddress)
	fmt.Println("Password: " + acc.Password)
	fmt.Println("Created At: " + acc.CreatedAt.String())

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lr *LoginRequest) Print() {
	fmt.Println("Username: " + lr.Username)
	fmt.Println("Password: " + lr.Password)
}
