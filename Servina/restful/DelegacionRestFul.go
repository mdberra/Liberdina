package restful

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type DelegacionRestFul struct {
}

func (rest *DelegacionRestFul) GetDelegacion(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GET")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	delegacion.CleanDelegacion()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["idDelegacion"])
	if err != nil {
		panic(err)
	}
	delegacion.ID_DELEGACION = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := delegacion.GetDelegacion(conexion.Db)
	if err != nil {
		log.Println("error GetDelegacion")
		panic(err)
	}

	j, err := json.Marshal(nAux)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}

func (rest *DelegacionRestFul) GetDelegaciones(w http.ResponseWriter, r *http.Request) {
	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Delegaciones, err := delegaciones.GetDelegaciones(conexion.Db)
	//	Delegaciones, err := delegacion.GetDelegaciones(conexion.Db)
	if err != nil {
		log.Println("error GetAllDelegacion")
		panic(err)
	}

	j, err := json.Marshal(Delegaciones)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *DelegacionRestFul) PostDelegacion(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un POST")
	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	delegacion.CleanDelegacion()

	//decodifica lo que viene y lo coloca en n
	err := json.NewDecoder(r.Body).Decode(&delegacion)
	if err != nil {
		panic(err)
	}

	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	delegacion.FECHA_ALTA = time.Now()

	if err := delegacion.CreateDelegacion(conexion.Db); err != nil {
		log.Println("error CreateCliente")
		panic(err)
	}

	j, err := json.Marshal(cliente)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *DelegacionRestFul) GetAllDelegacion(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GETAll")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Delegaciones, err := delegacion.GetDelegaciones(conexion.Db)
	if err != nil {
		log.Println("error GetAllCliente")
		panic(err)
	}

	j, err := json.Marshal(Delegaciones)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *DelegacionRestFul) PutDelegacion(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un PUT")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	delegacion.CleanDelegacion()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	_, err := delegacion.GetDelegacion(conexion.Db)
	if err != nil {
		// Delegacion no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		//		delegacion.TIPO_DOC = nAux.TIPO_DOC
		//		delegacion.NRO_DOC = nAux.NRO_DOC
		////////////////////////////
	}

	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&delegacion)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// n.FechaLastUpdate = time.Now()
	////////////////////////////
	cant, err := delegacion.UpdateDelegacion(conexion.Db)
	if err != nil {
		log.Println("error UpdateDelegacion")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	delegacion, _ = delegacion.GetDelegacion(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(delegacion)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}

func (rest *DelegacionRestFul) DeleteDelegacion(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un DELETE")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	delegacion.CleanDelegacion()

	//decodifica lo que viene y lo coloca en n
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	delegacion.ID_DELEGACION = int64(id)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := delegacion.DeleteDelegacion(conexion.Db)
	if err != nil {
		log.Println("error DeleteCliente")
		panic(err)
	} else {
		log.Println("DeleteCliente id: ", delegacion.ID_DELEGACION, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
