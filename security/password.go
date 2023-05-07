package security

import "golang.org/x/crypto/bcrypt" //O(1)

func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //O(n)
	if err != nil {                                                                  //O(1)
		return "", err //O(1)
	}
	return string(hashed), nil //O(1)
}

func VerifyPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) //O(n)
}
