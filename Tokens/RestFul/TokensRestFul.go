package RestFul

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type TokensRestFul struct {
}

func (rest *TokensRestFul) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuario.CleanUsuario()

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}
	if usuario.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Email no informado\n")
		return
	}
	if usuario.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Password no informada\n")
		return
	}
	//	spew.Dump(usuario)
	// Valido existencia del email
	if err := conexion.Open(conexion.GetEsquemaSeguridad()); err != nil {
		panic(err)
	} else {
		defer conexion.Close()
	}
	u, err := usuario.GetUsuarioEmail(conexion.Db)
	if err != nil {
		panic(err)
	}
	if u.IdUsuario > 0 { //ya existe
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("email %v ya existe", usuario.Email))
		return
	}
	//
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	usuario.Password = string(hashPassword)
	usuario.Fecha = time.Now().Add(time.Hour * -3)

	if err := usuario.CreateUsuario(conexion.Db); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(usuario)
}
func (rest *TokensRestFul) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuario.CleanUsuario()

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}
	if usuario.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Email no informado\n")
		return
	}
	if usuario.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Password no informada\n")
		return
	}
	if err := conexion.Open(conexion.GetEsquemaSeguridad()); err != nil {
		panic(err)
	} else {
		defer conexion.Close()
	}
	// Email no existe
	user, err := usuario.GetUsuarioEmail(conexion.Db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Email no existe")) //fmt.Errorf("user %q (id %d) not found", name, id))
		return
	}
	// comparo las passwords (DB hashed  con  ASCII ingresada)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(usuario.Password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Password Invalida\n")
		return
	}
	// esta todo bien
	token, err := usuario.GenerarToken()
	if err != nil {
		log.Fatal(err)
	}
	//	spew.Dump(token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
func (rest *TokensRestFul) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint protegido")
}
func (rest *TokensRestFul) ProtectedEndpoint1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint protegido1")
}
func (rest *TokensRestFul) ProtectedEndpoint2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint protegido2")
}

// https://github.com/gorilla/mux   Middleware
func (rest *TokensRestFul) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	// hacemos algo y luego llama al proximo Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "Token Invalido\n")
			return
		} else {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, funcion)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err.Error)
				return
			}
			//			spew.Dump(token)
			if token.Valid {
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err.Error)
				return
			}
		}
	})
}
func funcion(token *jwt.Token) (interface{}, error) {
	//	spew.Dump(token)
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("hubo un error")
	}
	return []byte(os.Getenv("Secreta")), nil
}
