package RestFul

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type RecetaRestFul struct {
}

func (rest *RecetaRestFul) GetReceta(w http.ResponseWriter, r *http.Request) {
	log.Println("GetReceta")

	receta.CleanReceta()
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["idReceta"])
	if err != nil {
		panic(err)
	}
	receta.IdReceta = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := receta.GetReceta(conexion.Db)
	if err != nil {
		log.Printf("RecetaRestFul.GetReceta: nAux %v", nAux)
		panic(err)
	}

	j, err := json.Marshal(nAux)
	//	var jsons []string  uint8
	//	for i := range j {
	//		jsons.append(i)
	//	}
	w.Write(j)
	//	log.Printf(jsons)
	/**	if err != nil {
	  		log.Println("problema en el json.Marshal")
	  	} else {
	  		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
	  		if nAux.IdReceta == -1 {
	  			w.WriteHeader(http.StatusNotFound) // 404
	  		} else {
	  //			jsons.append(jsons, String(j.))
	  			//
	  			// RecetaItems
	  			//
	  			nRItems, err := receta.GetRecetaItem(conexion.Db)
	  			if err != nil {
	  				log.Printf("RecetaRestFul.GetReceta: nRItems %v", nRItems)
	  				panic(err)
	  			}
	  			j, err = json.Marshal(nRItems)
	  			if err != nil {
	  				log.Println("problema en el json.Marshal")
	  			} else {
	  				w.Header().Set("Content-Type", "application/json") // standard del protocolo http
	  				w.WriteHeader(http.StatusOK)                       // escribe el 200
	  //				jsons.append(j)
	  			}
	  			//
	  			// RecetaFarmacia
	  			//
	  			nRFarmacia, err := receta.GetRecetaFarmacia(conexion.Db)
	  			if err != nil {
	  				log.Printf("RecetaRestFul.GetRecetaFarmacia: nRFarmacia %v", nRFarmacia)
	  				panic(err)
	  			}
	  			j, err := json.Marshal(nRFarmacia)
	  			if err != nil {
	  				log.Println("problema en el json.Marshal")
	  			} else {
	  //				jsons.append(j)

	  				w.Write(jsons)
	  			}
	  		}
	  	}*/
}
func (rest *RecetaRestFul) GetAllReceta(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllReceta")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	recetas, err := receta.GetRecetas(conexion.Db)
	if err != nil {
		log.Println("error GetAllReceta")
		panic(err)
	}

	j, err := json.Marshal(recetas)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *RecetaRestFul) PostReceta(w http.ResponseWriter, r *http.Request) {
	receta.CleanReceta()

	err := json.NewDecoder(r.Body).Decode(&receta)
	if err != nil {
		panic(err)
	}
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	receta.FechaCreacion = time.Now().Add(time.Hour * -3)
	receta.FechaEstado = time.Now().Add(time.Hour * -3)

	if err := receta.CreateReceta(conexion.Db); err != nil {
		log.Println("error CreateReceta")
		panic(err)
	}

	j, err := json.Marshal(receta)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *RecetaRestFul) PostRecetaItem(w http.ResponseWriter, r *http.Request) {
	// el decode directamente convierte en array
	err := json.NewDecoder(r.Body).Decode(&recetaItems)
	if err != nil {
		panic(err)
	}
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}

	if err := receta.CreateRecetaItem(conexion.Db, recetaItems); err != nil {
		log.Println("error CreateRecetaItems")
		panic(err)
	}

	j, err := json.Marshal(recetaItems)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *RecetaRestFul) PutReceta(w http.ResponseWriter, r *http.Request) {
	log.Println("PutReceta")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	receta.CleanReceta()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := receta.GetReceta(conexion.Db)
	if err != nil {
		// Receta no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		receta.FechaCreacion = nAux.FechaCreacion
		receta.IdMedico = nAux.IdMedico
		receta.Estado = nAux.Estado
		receta.FechaEstado = nAux.FechaEstado
		////////////////////////////
	}
	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&receta)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// receta.FechaLastUpdate = time.Now()
	receta.FechaEstado = time.Now().Add(time.Hour * -3)

	cant, err := receta.UpdateReceta(conexion.Db)
	if err != nil {
		log.Println("error UpdateReceta")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	receta, _ = receta.GetReceta(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(receta)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}
func (rest *RecetaRestFul) DeleteReceta(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteReceta")

	receta.CleanReceta()

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaReceDig()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := receta.DeleteReceta(conexion.Db)
	if err != nil {
		log.Println("error DeleteReceta")
		panic(err)
	} else {
		log.Println("DeleteReceta dni: ", receta.IdReceta, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
