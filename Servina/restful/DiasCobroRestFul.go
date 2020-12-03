package restful

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DiasCobroRestFul struct {
}

func (rest *DiasCobroRestFul) GetDiasCobro(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GET")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	diasCobro.CleanDiasCobro()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["idDelegacion"])
	if err != nil {
		panic(err)
	}
	diasCobro.ID_DELEGACION = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	DiasCobros, err := diasCobro.GetDiasCobroDelegacion(conexion.Db)
	if err != nil {
		log.Println("error GetDiasCobro")
		panic(err)
	}

	j, err := json.Marshal(DiasCobros)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}

func (rest *DiasCobroRestFul) PostDiasCobro(w http.ResponseWriter, r *http.Request) {
	diasCobro.CleanDiasCobro()

	//decodifica lo que viene y lo coloca en n
	err := json.NewDecoder(r.Body).Decode(&diasCobro)
	if err != nil {
		panic(err)
	}

	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}

	if err := diasCobro.CreateDiasCobro(conexion.Db); err != nil {
		log.Println("error CreateDiasCobro")
		panic(err)
	}

	j, err := json.Marshal(diasCobro)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}

func (rest *DiasCobroRestFul) DeleteDiasCobro(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un DELETE")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	diasCobro.CleanDiasCobro()

	//decodifica lo que viene y lo coloca en n
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["idDiasCobro"])
	if err != nil {
		panic(err)
	}
	diasCobro.ID_DIASCOBRO = int64(id)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := diasCobro.DeleteDiasCobro(conexion.Db)
	if err != nil {
		log.Println("error DeleteCliente")
		panic(err)
	} else {
		log.Println("DeleteCliente id: ", diasCobro.ID_DIASCOBRO, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
