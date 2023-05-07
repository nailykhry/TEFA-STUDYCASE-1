package util

import "errors" //O(1)

var (
	ErrInvalidEmail       = errors.New("invalid email")           //O(1)
	ErrEmailAlreadyExists = errors.New("email already exists")    //O(1)
	ErrEmptyPassword      = errors.New("password can't be empty") //O(1)
	ErrInvalidAuthToken   = errors.New("invalid auth-token")      //O(1)
	ErrInvalidCredentials = errors.New("invalid credentials")     //O(1)
	ErrUnauthorized       = errors.New("Unauthorized")            //O(1)
)
