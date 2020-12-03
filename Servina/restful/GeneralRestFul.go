package restful

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GeneralRestFul struct {
}

func (rest *GeneralRestFul) GetDebito(w http.ResponseWriter, r *http.Request) {
	debito.CleanDebito()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["idCliente"])
	if err != nil {
		panic(err)
	}
	debito.ID_CLIENTE = int64(id)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Debito, err := debito.GetDebito(conexion.Db)
	if err != nil {
		log.Println("error GetDebito")
		panic(err)
	}

	j, err := json.Marshal(Debito)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *GeneralRestFul) GetGeneral(w http.ResponseWriter, r *http.Request) {
	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	General, err := general.GetGeneral(conexion.Db)
	if err != nil {
		log.Println("error GetGeneral")
		panic(err)
	}

	j, err := json.Marshal(General)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *GeneralRestFul) GetFindCliente(w http.ResponseWriter, r *http.Request) {
	findCliente.CleanFindCliente()

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	FindCliente, err := findCliente.GetFindCliente(conexion.Db)
	if err != nil {
		log.Println("error findCliente.getFindCliente")
		panic(err)
	}

	j, err := json.Marshal(FindCliente)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}

/**
func (rest *GeneralRestFul) GetDelegaciones(w http.ResponseWriter, r *http.Request) {
	delegaciones.CleanDelegaciones()

	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Delegaciones, err := delegaciones.GetDelegaciones(conexion.Db)
	if err != nil {
		log.Println("error delegaciones.GetDelegaciones")
		panic(err)
	}

	j, err := json.Marshal(Delegaciones)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
*/
