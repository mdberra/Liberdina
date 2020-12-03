package RestFul

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type FarmaciaRestFul struct {
}

func (rest *FarmaciaRestFul) GetFarmacia(w http.ResponseWriter, r *http.Request) {
	log.Println("GetFarmacia")

	farmacia.CleanFarmacia()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["idFarmacia"])
	if err != nil {
		panic(err)
	}
	farmacia.IdFarmacia = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := farmacia.GetFarmacia(conexion.Db)
	if err != nil {
		log.Printf("FarmaciaRestFul.GetFarmacia: nAux %v", nAux)
		panic(err)
	}
	j, err := json.Marshal(nAux)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		if nAux.IdFarmacia == -1 {
			w.WriteHeader(http.StatusNotFound) // 404
		} else {
			w.WriteHeader(http.StatusOK) // escribe el 200
		}
		w.Write(j) // escribe el json
	}
}
func (rest *FarmaciaRestFul) GetAllFarmacia(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllFarmacia")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Farmacias, err := farmacia.GetFarmacias(conexion.Db)
	if err != nil {
		log.Println("error GetAllFarmacia")
		panic(err)
	}

	j, err := json.Marshal(Farmacias)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *FarmaciaRestFul) PostFarmacia(w http.ResponseWriter, r *http.Request) {
	farmacia.CleanFarmacia()

	err := json.NewDecoder(r.Body).Decode(&farmacia)
	if err != nil {
		panic(err)
	}
	log.Println("PostFarmacia: " + strconv.FormatInt(farmacia.IdFarmacia, 10))
	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	farmacia.FechaIngreso = time.Now().Add(time.Hour * -3)
	farmacia.FechaEstado = time.Now().Add(time.Hour * -3)

	if err := farmacia.CreateFarmacia(conexion.Db); err != nil {
		log.Println("error CreateFarmacia")
		panic(err)
	}

	j, err := json.Marshal(farmacia)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *FarmaciaRestFul) PutFarmacia(w http.ResponseWriter, r *http.Request) {
	log.Println("PutFarmacia")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	farmacia.CleanFarmacia()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := farmacia.GetFarmacia(conexion.Db)
	if err != nil {
		// Farmacia no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		farmacia.Nombre = nAux.Nombre
		farmacia.Direccion = nAux.Direccion
		farmacia.Telefono = nAux.Telefono
		farmacia.Email = nAux.Email
		farmacia.FechaIngreso = nAux.FechaIngreso
		farmacia.Estado = nAux.Estado
		farmacia.FechaEstado = nAux.FechaEstado
		////////////////////////////
	}

	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&farmacia)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// farmacia.FechaLastUpdate = time.Now()
	farmacia.FechaEstado = time.Now().Add(time.Hour * -3)

	cant, err := farmacia.UpdateFarmacia(conexion.Db)
	if err != nil {
		log.Println("error UpdateFarmacia")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	farmacia, _ = farmacia.GetFarmacia(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(farmacia)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}
func (rest *FarmaciaRestFul) DeleteFarmacia(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteFarmacia")

	farmacia.CleanFarmacia()

	farmacia.IdFarmacia = TransformRequestToDni(r)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := farmacia.DeleteFarmacia(conexion.Db)
	if err != nil {
		log.Println("error DeleteFarmacia")
		panic(err)
	} else {
		log.Println("DeleteFarmacia IdFarmacia: ", farmacia.IdFarmacia, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}