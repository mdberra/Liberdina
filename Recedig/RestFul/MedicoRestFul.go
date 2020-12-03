package RestFul

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MedicoRestFul struct {
}

func (rest *MedicoRestFul) GetMedico(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMedico")

	medico.CleanMedico()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["idMedico"])
	if err != nil {
		panic(err)
	}
	medico.IdMedico = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := medico.GetMedico(conexion.Db)
	if err != nil {
		log.Printf("MedicoRestFul.GetMedico: nAux %v", nAux)
		panic(err)
	}
	j, err := json.Marshal(nAux)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		if nAux.IdMedico == -1 {
			w.WriteHeader(http.StatusNotFound) // 404
		} else {
			w.WriteHeader(http.StatusOK) // escribe el 200
		}
		w.Write(j) // escribe el json
	}
}
func (rest *MedicoRestFul) GetMedicoDni(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMedicoDni")

	medico.CleanMedico()

	medico.Dni = TransformRequestToDni(r)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	medicos, err := medico.GetMedicosDni(conexion.Db)
	if err != nil {
		log.Println("error GetAllMedico")
		panic(err)
	}

	j, err := json.Marshal(medicos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicoRestFul) GetAllMedico(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllMedico")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	medicos, err := medico.GetMedicos(conexion.Db)
	if err != nil {
		log.Println("error GetAllMedico")
		panic(err)
	}

	j, err := json.Marshal(medicos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicoRestFul) PostMedico(w http.ResponseWriter, r *http.Request) {
	medico.CleanMedico()

	err := json.NewDecoder(r.Body).Decode(&medico)
	if err != nil {
		panic(err)
	}
	log.Println("PostMedico: " + strconv.FormatInt(medico.Dni, 10))
	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	medico.FechaIngreso = time.Now().Add(time.Hour * -3)
	medico.FechaEstado = time.Now().Add(time.Hour * -3)

	if err := medico.CreateMedico(conexion.Db); err != nil {
		log.Println("error CreateMedico")
		panic(err)
	}

	j, err := json.Marshal(medico)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicoRestFul) PutMedico(w http.ResponseWriter, r *http.Request) {
	log.Println("PutMedico")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	medico.CleanMedico()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := medico.GetMedico(conexion.Db)
	if err != nil {
		// Medico no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		medico.Nombre = nAux.Nombre
		medico.Apellido = nAux.Apellido
		medico.Email = nAux.Email
		medico.Telefono = nAux.Telefono
		medico.Dni = nAux.Dni
		medico.Matricula = nAux.Matricula
		medico.IdImagen = nAux.IdImagen
		medico.FechaIngreso = nAux.FechaIngreso
		medico.Estado = nAux.Estado
		medico.FechaEstado = nAux.FechaEstado
		////////////////////////////
	}

	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&medico)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// Medico.FechaLastUpdate = time.Now()
	medico.FechaEstado = time.Now().Add(time.Hour * -3)

	cant, err := medico.UpdateMedico(conexion.Db)
	if err != nil {
		log.Println("error UpdateMedico")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	medico, _ = medico.GetMedico(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(medico)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicoRestFul) DeleteMedico(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteMedico")

	medico.CleanMedico()

	medico.Dni = TransformRequestToDni(r)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := medico.DeleteMedico(conexion.Db)
	if err != nil {
		log.Println("error DeleteMedico")
		panic(err)
	} else {
		log.Println("DeleteMedico dni: ", medico.Dni, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
