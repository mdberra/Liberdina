package main

import (
	"log"
	"net/http"
	"time"

	restful "github.com/Liberdina/Servina/restful"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	clienteRest    restful.ClienteRestFul
	delegacionRest restful.DelegacionRestFul
	bancoRest      restful.BancoRestFul
	diasCobroRest  restful.DiasCobroRestFul
	generalRest    restful.GeneralRestFul
	contactoRest   restful.ContactoRestFul
	imageRest      restful.ImageRestFul
	ping           restful.Ping
)

// paquete mux es un enrutador avanzado que permite hacer el CRUD con RestFul
func main() {
	ping.PingDB()
	//	ping.PingStorage()

	router := mux.NewRouter()
	router.StrictSlash(false) //permite que estas dos rutas sean iguales   /api/user y /api/user/

	// General
	router.HandleFunc("/api/keyvalue/{entidad:[a-z]+}", contactoRest.GetKeyValue).Methods("GET")
	router.HandleFunc("/api/findCliente", generalRest.GetFindCliente).Methods("GET")         //ReadAll
	router.HandleFunc("/api/cliente/{nroDoc:[0-9]+}", clienteRest.GetCliente).Methods("GET") //Read
	router.HandleFunc("/api/clientes", clienteRest.GetAllCliente).Methods("GET")             //ReadAll
	router.HandleFunc("/api/cliente", clienteRest.PostCliente).Methods("POST")               //Create
	router.HandleFunc("/api/cliente", clienteRest.PutCliente).Methods("PUT")                 //Update
	//	router.HandleFunc("/api/cliente/{id:[0-9]+}", clienteRest.DeleteCliente).Methods("DELETE") //Delete

	// PepeYa
	router.HandleFunc("/api/pepeya/pingDB", ping.GetPingDB).Methods("GET")
	router.HandleFunc("/api/pepeya/pingStorage", ping.GetPingStorage).Methods("GET")

	router.HandleFunc("/api/pepeya/contactoDni/{dni:[0-9]+}", contactoRest.GetContactoDni).Methods("GET")  //Read
	router.HandleFunc("/api/pepeya/contacto/{idContacto:[0-9]+}", contactoRest.GetContacto).Methods("GET") //Read
	router.HandleFunc("/api/pepeya/contactos", contactoRest.GetAllContacto).Methods("GET")                 //ReadAll
	router.HandleFunc("/api/pepeya/contacto", contactoRest.PostContacto).Methods("POST")                   //Create
	router.HandleFunc("/api/pepeya/contacto", contactoRest.PutContacto).Methods("PUT")
	//	router.HandleFunc("/api/pepeya/contacto/{dni:[0-9]+}", contactoRest.DeleteContacto).Methods("DELETE")  //Delete
	router.HandleFunc("/api/pepeya/image/{fileName}", imageRest.GetImage).Methods("GET")
	router.HandleFunc("/api/pepeya/image", imageRest.PostImage).Methods("POST")

	// Servina
	router.HandleFunc("/api/servina/debitos/{idCliente:[0-9]+}", generalRest.GetDebito).Methods("GET")
	router.HandleFunc("/api/servina/estadisticas", generalRest.GetGeneral).Methods("GET")

	router.HandleFunc("/api/servina/delegaciones", delegacionRest.GetDelegaciones).Methods("GET")
	router.HandleFunc("/api/servina/delegacion/{idDelegacion:[0-9]+}", delegacionRest.GetDelegacion).Methods("GET")
	router.HandleFunc("/api/servina/delegacion", delegacionRest.PostDelegacion).Methods("POST")
	router.HandleFunc("/api/servina/delegacion", delegacionRest.PutDelegacion).Methods("PUT")
	router.HandleFunc("/api/servina/delegacion/{dni:[0-9]+}", delegacionRest.DeleteDelegacion).Methods("DELETE")

	router.HandleFunc("/api/servina/bancos", bancoRest.GetBancos).Methods("GET")
	router.HandleFunc("/api/servina/banco/{idBanco:[0-9]+}", bancoRest.GetBanco).Methods("GET")
	router.HandleFunc("/api/servina/banco", bancoRest.PostBanco).Methods("POST")
	router.HandleFunc("/api/servina/banco", bancoRest.PutBanco).Methods("PUT")
	router.HandleFunc("/api/servina/banco/{id:[0-9]+}", bancoRest.DeleteBanco).Methods("DELETE")

	router.HandleFunc("/api/servina/diasCobro/{idDelegacion:[0-9]+}", diasCobroRest.GetDiasCobro).Methods("GET")
	router.HandleFunc("/api/servina/diasCobro", diasCobroRest.PostDiasCobro).Methods("POST")
	router.HandleFunc("/api/servina/diasCobro/{idDiasCobro:[0-9]+}", diasCobroRest.DeleteDiasCobro).Methods("DELETE")

	corsOpts := cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				//			http.MethodPatch,
				//			http.MethodOptions,
				http.MethodHead,
			},
			AllowedHeaders: []string{
				"*", //or you can your header key values which you are using in your application
			},
		},
	)
	server := &http.Server{
		Addr: ":5000",
		//		Handler:        cors.Default().Handler(router),
		Handler:        corsOpts.Handler(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1Mb
	}
	log.Println("Escuchando en :5000")
	server.ListenAndServe()
}
