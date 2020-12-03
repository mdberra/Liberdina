package RestFul

import (
	"log"
	"net/http"
)

type Ping struct {
}

func (rest *Ping) GetPingDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msg := []byte("")
	if err := conexion.ConectToDB(); err != nil {
		msg = []byte("Cloud Mysql - Conexion Error " + err.Error())
		w.WriteHeader(http.StatusRequestTimeout)
	} else {
		msg = []byte("Cloud Mysql - Conexion Exitosa")
		w.WriteHeader(http.StatusOK)
	}
	log.Println(msg)
	w.Write(msg)
}
func (rest *Ping) GetPingStorage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msg := []byte("")

	if err := cloudStorageService.Connect(); err != nil {
		msg = []byte("CloudStorage - Conexion Error " + err.Error())
		w.WriteHeader(http.StatusRequestTimeout)
	} else {
		//		client := cloudStorageService.GetClient()
		//		out, _ := json.Marshal(client)
		msg = []byte("CloudStorage - Conexion Exitosa") // + string(out))
		w.WriteHeader(http.StatusOK)
	}
	log.Println(string(msg))
	w.Write(msg)
}

func (rest *Ping) PingDB() {
	if err := conexion.ConectToDB(); err != nil {
		log.Println(err)
	} else {
		log.Println("Cloud Mysql - Conexion Exitosa")
	}
}

func (rest *Ping) PingStorage() {
	if err := cloudStorageService.Connect(); err != nil {
		log.Fatalf("No se puede contectar: %v", err)
	} else {
		log.Println("CloudStorage - Conexion Exitosa")
	}
}
