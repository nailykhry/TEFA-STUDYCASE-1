package util

import "strings"

type JError struct {
	Error string `json:"error"` //O(1)
}

func NewJError(err error) JError {
	jerr := JError{"generic error"} //O(1)
	if err != nil {                 //O(1)
		jerr.Error = err.Error() //O(1)
	}
	return jerr //O(1)
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email)) //O(1)
}
