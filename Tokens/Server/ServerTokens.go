package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Tokens/Data"
	"github.com/Tokens/RestFul"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

var (
	conexion   Data.Conexion
	ping       RestFul.Ping
	tokensRest RestFul.TokensRestFul
)

func init() {
	gotenv.Load()
}
func main() {
	conexion.ConectToDB()

	router := mux.NewRouter()
	router.StrictSlash(false) //permite que estas dos rutas sean iguales   /api/user y /api/user/

	router.HandleFunc("/api/tokens/pingDB", ping.PingDB).Methods("GET")

	router.HandleFunc("/api/tokens/signup", tokensRest.Signup).Methods("POST")
	router.HandleFunc("/api/tokens/login", tokensRest.Login).Methods("POST")
	router.HandleFunc("/api/tokens/protected", tokensRest.TokenVerifyMiddleWare(tokensRest.ProtectedEndpoint)).Methods("GET")
	router.HandleFunc("/api/tokens/protected1", tokensRest.TokenVerifyMiddleWare(tokensRest.ProtectedEndpoint1)).Methods("GET")
	router.HandleFunc("/api/tokens/protected2", tokensRest.TokenVerifyMiddleWare(tokensRest.ProtectedEndpoint2)).Methods("GET")

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	//Server
	server := &http.Server{
		Addr:           ":5200",
		Handler:        cors.Default().Handler(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1Mb
	}
	log.Println("Escuchando en :5200")
	server.ListenAndServe()
}
