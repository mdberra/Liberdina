package restful

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ClienteRestFul struct {
}

func (rest *ClienteRestFul) GetCliente(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GET")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	cliente.CleanCliente()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["nroDoc"])
	if err != nil {
		panic(err)
	}
	cliente.NRO_DOC = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := cliente.GetCliente(conexion.Db)
	if err != nil {
		log.Println("error GetCliente")
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

func (rest *ClienteRestFul) GetEstad(w http.ResponseWriter, r *http.Request) {
	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Clientes, err := cliente.GetClientes(conexion.Db)
	if err != nil {
		log.Println("error GetAllCliente")
		panic(err)
	}

	j, err := json.Marshal(Clientes)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}

func (rest *ClienteRestFul) PostCliente(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un POST")
	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	cliente.CleanCliente()

	//decodifica lo que viene y lo coloca en n
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		panic(err)
	}
	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	cliente.F_INGRESO = time.Now()
	cliente.F_ESTADO = time.Now()

	if err := cliente.CreateCliente(conexion.Db); err != nil {
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
func (rest *ClienteRestFul) GetAllCliente(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GETAll")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Clientes, err := cliente.GetClientes(conexion.Db)
	if err != nil {
		log.Println("error GetAllCliente")
		panic(err)
	}

	j, err := json.Marshal(Clientes)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *ClienteRestFul) PutCliente(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un PUT")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	cliente.CleanCliente()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := cliente.GetCliente(conexion.Db)
	if err != nil {
		// Cliente no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		cliente.TIPO_DOC = nAux.TIPO_DOC
		cliente.NRO_DOC = nAux.NRO_DOC
		////////////////////////////
	}

	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// n.FechaLastUpdate = time.Now()
	////////////////////////////
	cant, err := cliente.UpdateCliente(conexion.Db)
	if err != nil {
		log.Println("error UpdateCliente")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	cliente, _ = cliente.GetCliente(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(cliente)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}

func (rest *ClienteRestFul) DeleteCliente(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un DELETE")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	cliente.CleanCliente()

	//decodifica lo que viene y lo coloca en n
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	cliente.ID_CLIENTE = int64(id)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := cliente.DeleteCliente(conexion.Db)
	if err != nil {
		log.Println("error DeleteCliente")
		panic(err)
	} else {
		log.Println("DeleteCliente id: ", cliente.ID_CLIENTE, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
