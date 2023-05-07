package security

import (
	"TEFA-STUDYCASE-1/util" //O(1)
	"fmt"                   //O(1)
	"os"                    //O(1)
	"time"                  //O(1)

	jwt "github.com/form3tech-oss/jwt-go" //O(1)
)

var (
	JwtSecretKey     = []byte(os.Getenv("JWT_SECRET_KEY")) //O(1)
	JwtSigningMethod = jwt.SigningMethodHS256.Name         //O(1)
)

func NewToken(userId string, userRole string) (string, error) {
	claims := jwt.StandardClaims{ //O(1)
		Id:        userId,
		Subject:   userRole,
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //O(n)
	return token.SignedString(JwtSecretKey)                    //O(1)
}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //O(1)
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) //O(1)
	}
	return JwtSecretKey, nil //O(1)
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := new(jwt.StandardClaims)                                            //O(1)
	token, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod) //O(n)
	if err != nil {                                                              //O(1)
		return nil, err //O(1)
	}
	var ok bool                                     //O(1)
	claims, ok = token.Claims.(*jwt.StandardClaims) //O(1)
	if !ok || !token.Valid {                        //O(1)
		return nil, util.ErrInvalidAuthToken //O(1)
	}
	return claims, nil //O(1)
}
