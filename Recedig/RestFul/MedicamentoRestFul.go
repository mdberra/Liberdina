package RestFul

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MedicamentoRestFul struct {
}

func (rest *MedicamentoRestFul) GetMedicamento(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMedicamento")

	medicamento.CleanMedicamento()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	medicamento.Id = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := medicamento.GetMedicamento(conexion.Db)
	if err != nil {
		log.Printf("MedicamentoRestFul.GetMedicamento: nAux %v", nAux)
		panic(err)
	}
	j, err := json.Marshal(nAux)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		if nAux.Id == -1 {
			w.WriteHeader(http.StatusNotFound) // 404
		} else {
			w.WriteHeader(http.StatusOK) // escribe el 200
		}
		w.Write(j) // escribe el json
	}
}
func (rest *MedicamentoRestFul) GetAllMedicamento(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllMedicamento")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Medicamentos, err := medicamento.GetMedicamentos(conexion.Db)
	if err != nil {
		log.Println("error GetAllMedicamento")
		panic(err)
	}

	j, err := json.Marshal(Medicamentos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicamentoRestFul) PostMedicamento(w http.ResponseWriter, r *http.Request) {
	medicamento.CleanMedicamento()

	err := json.NewDecoder(r.Body).Decode(&medicamento)
	if err != nil {
		panic(err)
	}
	log.Println("PostMedicamento: " + strconv.FormatInt(medicamento.Id, 10))
	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	medicamento.FechaIngreso = time.Now().Add(time.Hour * -3)
	medicamento.FechaEstado = time.Now().Add(time.Hour * -3)

	if err := medicamento.CreateMedicamento(conexion.Db); err != nil {
		log.Println("error CreateMedicamento")
		panic(err)
	}

	j, err := json.Marshal(medicamento)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicamentoRestFul) PutMedicamento(w http.ResponseWriter, r *http.Request) {
	log.Println("PutMedicamento")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	medicamento.CleanMedicamento()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := medicamento.GetMedicamento(conexion.Db)
	if err != nil {
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		medicamento.Nombre = nAux.Nombre
		medicamento.Droga = nAux.Droga
		medicamento.IdLaboratorio = nAux.IdLaboratorio
		medicamento.FechaIngreso = nAux.FechaIngreso
		medicamento.Estado = nAux.Estado
		medicamento.FechaEstado = nAux.FechaEstado
		////////////////////////////
	}
	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&medicamento)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// medicamento.FechaLastUpdate = time.Now()
	medicamento.FechaEstado = time.Now().Add(time.Hour * -3)

	cant, err := medicamento.UpdateMedicamento(conexion.Db)
	if err != nil {
		log.Println("error UpdateMedicamento")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	medicamento, _ = medicamento.GetMedicamento(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(medicamento)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}
func (rest *MedicamentoRestFul) DeleteMedicamento(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteMedicamento")

	medicamento.CleanMedicamento()

	medicamento.Id = TransformRequestToDni(r)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := medicamento.DeleteMedicamento(conexion.Db)
	if err != nil {
		log.Println("error DeleteMedicamento")
		panic(err)
	} else {
		log.Println("DeleteMedicamento IdMedicamento: ", medicamento.Id, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
