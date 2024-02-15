package jwttoken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JoseGaldamez/kamehamehapi/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type (
	Token struct {
		Token string `json:"token"`
	}
	TokenRequest struct {
		Email string `json:"email"`
	}
	ResquestError struct {
		Error string `json:"error"`
	}
)

func CreateRouter(path string, router *mux.Router) {
	router.HandleFunc(utils.ApiUrlPrefix+path, createToken).Methods("POST")
	router.HandleFunc(utils.ApiUrlPrefix+path, getTokenInformation).Methods("GET")
}

func createToken(response http.ResponseWriter, request *http.Request) {
	var tokenRequest TokenRequest
	response.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(request.Body).Decode(&tokenRequest)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResquestError{Error: "Invalid parameters"})
		return
	}
	if tokenRequest.Email == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResquestError{Error: "Email is requerid"})
		return
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": tokenRequest.Email,
			"exp":   time.Now().Add(time.Hour * 24 * 365).Unix(),
		})

	tokenString, err := newToken.SignedString(utils.SecretKeyToken)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResquestError{Error: "Something went wrong creating token"})
		return
	}

	token := Token{Token: tokenString}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(token)
}

func getTokenInformation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	email, err := VerifyToken(request)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(ResquestError{Error: err.Error()})
		return
	}

	if email == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResquestError{Error: "Invalid header parameters."})
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{"email": email})
}

func VerifyToken(request *http.Request) (string, error) {

	var headerToken = request.Header.Get("Authorization")

	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		return utils.SecretKeyToken, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	var email string

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		email = fmt.Sprint(claims["email"])
	} else {
		return "", fmt.Errorf("invalid token")
	}

	return email, nil
}
