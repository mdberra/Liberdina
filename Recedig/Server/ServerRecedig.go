package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Liberdina/Recedig/RestFul"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	ping      RestFul.Ping
	imageRest RestFul.ImageRestFul

	medicoRest      RestFul.MedicoRestFul
	farmaciaRest    RestFul.FarmaciaRestFul
	laboratorioRest RestFul.LaboratorioRestFul
	medicamentoRest RestFul.MedicamentoRestFul
	recetaRest      RestFul.RecetaRestFul
)

// paquete mux es un enrutador avanzado que permite hacer el CRUD con RestFul
func main() {
	ping.PingDB()
	ping.PingStorage()

	router := mux.NewRouter()
	router.StrictSlash(false) //permite que estas dos rutas sean iguales   /api/user y /api/user/

	router.HandleFunc("/api/recedig/pingDB", ping.GetPingDB).Methods("GET")
	router.HandleFunc("/api/recedig/pingStorage", ping.GetPingStorage).Methods("GET")

	router.HandleFunc("/api/recedig/image/{fileName}", imageRest.GetImage).Methods("GET")
	router.HandleFunc("/api/recedig/image", imageRest.PostImage).Methods("POST")

	//	router.HandleFunc("/api/recedig/keyvalue", medicoRest.GetKeyValue).Methods("GET")

	router.HandleFunc("/api/recedig/medico/{idMedico:[0-9]+}", medicoRest.GetMedico).Methods("GET")  //Read
	router.HandleFunc("/api/recedig/medicoDni/{dni:[0-9]+}", medicoRest.GetMedicoDni).Methods("GET") //Read
	router.HandleFunc("/api/recedig/medicos", medicoRest.GetAllMedico).Methods("GET")                //ReadAll
	router.HandleFunc("/api/recedig/medico", medicoRest.PostMedico).Methods("POST")                  //Create
	router.HandleFunc("/api/recedig/medico", medicoRest.PutMedico).Methods("PUT")                    //Update
	router.HandleFunc("/api/recedig/medico/{dni:[0-9]+}", medicoRest.DeleteMedico).Methods("DELETE") //Delete

	router.HandleFunc("/api/recedig/farmacia/{idFarmacia:[0-9]+}", farmaciaRest.GetFarmacia).Methods("GET")       //Read
	router.HandleFunc("/api/recedig/farmacias", farmaciaRest.GetAllFarmacia).Methods("GET")                       //ReadAll
	router.HandleFunc("/api/recedig/farmacia", farmaciaRest.PostFarmacia).Methods("POST")                         //Create
	router.HandleFunc("/api/recedig/farmacia", farmaciaRest.PutFarmacia).Methods("PUT")                           //Update
	router.HandleFunc("/api/recedig/farmacia/{idFarmacia:[0-9]+}", farmaciaRest.DeleteFarmacia).Methods("DELETE") //Delete

	router.HandleFunc("/api/recedig/laboratorio/{idLaboratorio:[0-9]+}", laboratorioRest.GetLaboratorio).Methods("GET")       //Read
	router.HandleFunc("/api/recedig/laboratorios", laboratorioRest.GetAllLaboratorio).Methods("GET")                          //ReadAll
	router.HandleFunc("/api/recedig/laboratorio", laboratorioRest.PostLaboratorio).Methods("POST")                            //Create
	router.HandleFunc("/api/recedig/laboratorio", laboratorioRest.PutLaboratorio).Methods("PUT")                              //Update
	router.HandleFunc("/api/recedig/laboratorio/{idLaboratorio:[0-9]+}", laboratorioRest.DeleteLaboratorio).Methods("DELETE") //Delete

	router.HandleFunc("/api/recedig/medicamento/{idMedicamento:[0-9]+}", medicamentoRest.GetMedicamento).Methods("GET")       //Read
	router.HandleFunc("/api/recedig/medicamentos", medicamentoRest.GetAllMedicamento).Methods("GET")                          //ReadAll
	router.HandleFunc("/api/recedig/medicamento", medicamentoRest.PostMedicamento).Methods("POST")                            //Create
	router.HandleFunc("/api/recedig/medicamento", medicamentoRest.PutMedicamento).Methods("PUT")                              //Update
	router.HandleFunc("/api/recedig/medicamento/{idMedicamento:[0-9]+}", medicamentoRest.DeleteMedicamento).Methods("DELETE") //Delete

	router.HandleFunc("/api/recedig/receta/{idReceta:[0-9]+}", recetaRest.GetReceta).Methods("GET") //Read
	router.HandleFunc("/api/recedig/recetas", recetaRest.GetAllReceta).Methods("GET")               //ReadAll
	router.HandleFunc("/api/recedig/receta", recetaRest.PostReceta).Methods("POST")
	router.HandleFunc("/api/recedig/recetaItem", recetaRest.PostRecetaItem).Methods("POST")
	router.HandleFunc("/api/recedig/receta", recetaRest.PutReceta).Methods("PUT")                         //Update
	router.HandleFunc("/api/recedig/receta/{idReceta:[0-9]+}", recetaRest.DeleteReceta).Methods("DELETE") //Delete

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	//Server
	server := &http.Server{
		Addr:           ":5100",
		Handler:        cors.Default().Handler(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1Mb
	}
	log.Println("Escuchando en :5100")
	server.ListenAndServe()
}
