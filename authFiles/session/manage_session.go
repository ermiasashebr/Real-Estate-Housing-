package session

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"github.com/ermiasashebr/housing/authFiles/csrfToken"
	"net/http"
	"time"
)

func Create(claims jwt.Claims, sessionID string, signingKey []byte, w http.ResponseWriter) {
	signedString, err := csrfToken.Generate(signingKey, claims)
	log.Println("Session signedstring", signedString)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	c := http.Cookie{
		Name:     sessionID,
		Value:    signedString,
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

func Valid(cookieValue string, signingKey []byte) (bool, error) {
	valid, err := csrfToken.Valid(cookieValue, signingKey)
	if err != nil || !valid {
		return false, errors.New("Invalid Session Cookie")
	}
	return true, nil
}

func Remove(sessionID string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:    sessionID,
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
		Value:   "",
	}
	http.SetCookie(w, &c)
}


