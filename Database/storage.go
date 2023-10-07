package database

import types "github.com/Arch-4ng3l/Bank/Types"

type Storage interface {
	SignUp(req *types.SignUpRequest) *types.Account
	Login(req *types.LoginRequest, encryptedPw string) *types.Account
	Transaction(tr *types.Transaction)
}
