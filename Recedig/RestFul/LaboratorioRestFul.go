package RestFul

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type LaboratorioRestFul struct {
}

func (rest *LaboratorioRestFul) GetLaboratorio(w http.ResponseWriter, r *http.Request) {
	log.Println("GetLaboratorio")

	laboratorio.CleanLaboratorio()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["idLaboratorio"])
	if err != nil {
		panic(err)
	}
	laboratorio.IdLaboratorio = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := laboratorio.GetLaboratorio(conexion.Db)
	if err != nil {
		log.Printf("LaboratorioRestFul.GetLaboratorio: nAux %v", nAux)
		panic(err)
	}
	j, err := json.Marshal(nAux)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		if nAux.IdLaboratorio == -1 {
			w.WriteHeader(http.StatusNotFound) // 404
		} else {
			w.WriteHeader(http.StatusOK) // escribe el 200
		}
		w.Write(j) // escribe el json
	}
}
func (rest *LaboratorioRestFul) GetAllLaboratorio(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllLaboratorio")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	laboratorios, err := laboratorio.GetLaboratorios(conexion.Db)
	if err != nil {
		log.Println("error GetAllLaboratorio")
		panic(err)
	}

	j, err := json.Marshal(laboratorios)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *LaboratorioRestFul) PostLaboratorio(w http.ResponseWriter, r *http.Request) {
	laboratorio.CleanLaboratorio()

	err := json.NewDecoder(r.Body).Decode(&laboratorio)
	if err != nil {
		panic(err)
	}
	log.Println("PostLaboratorio: " + strconv.FormatInt(laboratorio.IdLaboratorio, 10))
	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	laboratorio.FechaIngreso = time.Now().Add(time.Hour * -3)
	laboratorio.FechaEstado = time.Now().Add(time.Hour * -3)

	if err := laboratorio.CreateLaboratorio(conexion.Db); err != nil {
		log.Println("error CreateLaboratorio")
		panic(err)
	}

	j, err := json.Marshal(laboratorio)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *LaboratorioRestFul) PutLaboratorio(w http.ResponseWriter, r *http.Request) {
	log.Println("PutLaboratorio")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	laboratorio.CleanLaboratorio()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := laboratorio.GetLaboratorio(conexion.Db)
	if err != nil {
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		laboratorio.Nombre = nAux.Nombre
		laboratorio.Direccion = nAux.Direccion
		laboratorio.Telefono = nAux.Telefono
		laboratorio.Email = nAux.Email
		laboratorio.FechaIngreso = nAux.FechaIngreso
		laboratorio.Estado = nAux.Estado
		laboratorio.FechaEstado = nAux.FechaEstado
		////////////////////////////
	}

	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&laboratorio)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// laboratorio.FechaLastUpdate = time.Now()
	laboratorio.FechaEstado = time.Now().Add(time.Hour * -3)

	cant, err := laboratorio.UpdateLaboratorio(conexion.Db)
	if err != nil {
		log.Println("error UpdateLaboratorio")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	laboratorio, _ = laboratorio.GetLaboratorio(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(laboratorio)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}
func (rest *LaboratorioRestFul) DeleteLaboratorio(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteLaboratorio")

	laboratorio.CleanLaboratorio()

	laboratorio.IdLaboratorio = TransformRequestToDni(r)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := laboratorio.DeleteLaboratorio(conexion.Db)
	if err != nil {
		log.Println("error DeleteLaboratorio")
		panic(err)
	} else {
		log.Println("DeleteLaboratorio IdLaboratorio: ", laboratorio.IdLaboratorio, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
