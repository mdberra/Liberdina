package restful

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BancoRestFul struct {
}

func (rest *BancoRestFul) GetBanco(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GET")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	banco.CleanBanco()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["idBanco"])
	if err != nil {
		panic(err)
	}
	banco.ID_BANCO = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := banco.GetBanco(conexion.Db)
	if err != nil {
		log.Println("error GetBanco")
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

func (rest *BancoRestFul) GetBancos(w http.ResponseWriter, r *http.Request) {
	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Bancos, err := banco.GetBancos(conexion.Db)
	if err != nil {
		log.Println("error GetAllBanco")
		panic(err)
	}

	j, err := json.Marshal(Bancos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}

func (rest *BancoRestFul) PostBanco(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un POST")
	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	banco.CleanBanco()

	//decodifica lo que viene y lo coloca en n
	err := json.NewDecoder(r.Body).Decode(&banco)
	if err != nil {
		panic(err)
	}

	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}

	if err := banco.CreateBanco(conexion.Db); err != nil {
		log.Println("error CreateBanco")
		panic(err)
	}

	j, err := json.Marshal(banco)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *BancoRestFul) GetAllBanco(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un GETAll")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Bancos, err := banco.GetBancos(conexion.Db)
	if err != nil {
		log.Println("error GetAllCliente")
		panic(err)
	}

	j, err := json.Marshal(Bancos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *BancoRestFul) PutBanco(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un PUT")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	banco.CleanBanco()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	_, err := banco.GetBanco(conexion.Db)
	if err != nil {
		// Delegacion no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		//		banco.TIPO_DOC = nAux.TIPO_DOC
		//		banco.NRO_DOC = nAux.NRO_DOC
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
	cant, err := banco.UpdateBanco(conexion.Db)
	if err != nil {
		log.Println("error UpdateDelegacion")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	banco, _ = banco.GetBanco(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(delegacion)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}

func (rest *BancoRestFul) DeleteBanco(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hice un DELETE")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	banco.CleanBanco()

	//decodifica lo que viene y lo coloca en n
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	banco.ID_BANCO = int64(id)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaFinanciera()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := banco.DeleteBanco(conexion.Db)
	if err != nil {
		log.Println("error DeleteCliente")
		panic(err)
	} else {
		log.Println("DeleteCliente id: ", banco.ID_BANCO, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
