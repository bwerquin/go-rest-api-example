package helpers

import (   
	"encoding/json"  
	_ "github.com/gorilla/mux"
	"net/http"
	"strings"
	"log"
	"fmt"
	"crypto/rsa"
    "crypto/x509"
    "encoding/base64"

    jwt "github.com/dgrijalva/jwt-go"
	)

	var(
		hostname string
		realm string
		public_key_location string
		public_key_string string
		public_key *rsa.PublicKey
		err error
	)

	func InitializeOauthPublicKey(){ 
		hostname = AppConfig.HOST
		realm = AppConfig.REALM
		public_key_location = fmt.Sprintf("%s/realms/%s", hostname , realm)
		// Add HTTP GET to get back "public_key" value from JSON response
		public_key_string = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhYFctTXbfogKXa7tHvHMznidtHTiZ5wyaPF87xCNIBVKwCJJLNWCB3wjUu6V/N3CREG7BJwt+QCmuW48ww1l1PVuT5H7I83c8BBpUk49IAh6wPcsxkgA7hpbERE1dQV8huGn/HH5ONaIvAMzGStg+GrV9nksEnmDYZqtBVY/Y10uWIeiANEh64sEQb4zxhHGAd3D9ZIVNRZGeLRZgXS2azUvVsZwPjwEKs1hzR2SLl+PbPGbCzIugRAe96QAI4mOYnuzLelsz7v9u1Yve0yI8akqsdVuyOuLc+f3m0e/JkVi7voHxh3Ryom2vI6BA5kfaT+Oszn98t7E4vngD5V1uQIDAQAB"
		public_key, err = parseKeycloakRSAPublicKey(public_key_string)
		log.Println("Public Key parsed")
	}

	func parseKeycloakRSAPublicKey(base64Encoded string) (*rsa.PublicKey, error) {
		buf, err := base64.StdEncoding.DecodeString(base64Encoded)
		if err != nil {
			return nil, err
		}
		parsedKey, err := x509.ParsePKIXPublicKey(buf)
		if err != nil {
			return nil, err
		}
		publicKey, ok := parsedKey.(*rsa.PublicKey)
		if ok {
			return publicKey, nil
		}
		return nil, fmt.Errorf("unexpected key type %T", publicKey)
	}

	func Protect(next http.Handler) http.Handler {   
		return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {   
			
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) < 1 {
				w.WriteHeader(401)
				log.Println(fmt.Sprintf("UnauthorizedError, no header or len(authHeader) < 1 : %s", authHeader))
				json.NewEncoder(w).Encode(UnauthorizedError())
				return
			}

			token_string := strings.Split(authHeader," ")[1]
			log.Println(fmt.Sprintf("Check oidc token : %s", token_string))

			token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// return the public key that is used to validate the token.
				return public_key, nil
			})

			if err != nil{
				w.WriteHeader(400)
				log.Println(fmt.Sprintf("Error parsing or validating token: %s", err.Error()))
				json.NewEncoder(w).Encode(BadRequestError(err.Error()))
				return
			}
			
			if !token.Valid {
				w.WriteHeader(401)
				log.Println(fmt.Sprintf("UnauthorizedError, !isTokenValid = true - Header: %s ", authHeader))
				json.NewEncoder(w).Encode(UnauthorizedError())
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
                username := claims["preferred_username"].(string)
				log.Println(fmt.Sprintf("Username: %s", username))
          	}
			next.ServeHTTP(w, r)
			
		})
	}