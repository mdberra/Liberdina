package restful

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ContactoRestFul struct {
}

func (rest *ContactoRestFul) GetKeyValue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	keyValue.Entidad = vars["entidad"]

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Keyvalues, err := keyValue.GetKeyValue(conexion.Db)
	if err != nil {
		log.Println("error GetKeyValue")
		panic(err)
	}

	j, err := json.Marshal(Keyvalues)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *ContactoRestFul) GetContacto(w http.ResponseWriter, r *http.Request) {
	log.Println("GetContacto")

	contacto.CleanContacto()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["idContacto"])
	if err != nil {
		panic(err)
	}
	contacto.IdContacto = int64(id)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	nAux, err := contacto.GetContacto(conexion.Db)
	if err != nil {
		log.Printf("ContactoRestFul.GetContacto: nAux %v", nAux)
		panic(err)
	}
	j, err := json.Marshal(nAux)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		if nAux.IdContacto == -1 {
			w.WriteHeader(http.StatusNotFound) // 404
		} else {
			w.WriteHeader(http.StatusOK) // escribe el 200
		}
		w.Write(j) // escribe el json
	}
}
func (rest *ContactoRestFul) GetContactoDni(w http.ResponseWriter, r *http.Request) {
	log.Println("GetContactoDni")

	contacto.CleanContacto()

	contacto.Dni = TransformRequestToDni(r)

	//  BD Conexion y GET
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	Contactos, err := contacto.GetContactosDni(conexion.Db)
	if err != nil {
		log.Println("error GetAllContacto")
		panic(err)
	}

	j, err := json.Marshal(Contactos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *ContactoRestFul) GetAllContacto(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllContacto")

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// GETAll
	Contactos, err := contacto.GetContactos(conexion.Db)
	if err != nil {
		log.Println("error GetAllContacto")
		panic(err)
	}

	j, err := json.Marshal(Contactos)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200
		w.Write(j)                                         // escribe el json
	}
}
func (rest *ContactoRestFul) PostContacto(w http.ResponseWriter, r *http.Request) {
	contacto.CleanContacto()

	err := json.NewDecoder(r.Body).Decode(&contacto)
	if err != nil {
		panic(err)
	}
	log.Println("PostContacto: " + strconv.FormatInt(contacto.Dni, 10))
	//  BD Conexion e Insert
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	contacto.FechaIngreso = time.Now().Add(time.Hour * -3)
	contacto.FechaEstado = time.Now().Add(time.Hour * -3)

	if err := contacto.CreateContacto(conexion.Db); err != nil {
		log.Println("error CreateContacto")
		panic(err)
	}

	j, err := json.Marshal(contacto)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (rest *ContactoRestFul) PutContacto(w http.ResponseWriter, r *http.Request) {
	log.Println("PutContacto")

	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	contacto.CleanContacto()

	//  BD Conexion y UPDATE
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	// leo para ver que existe el registro y para completar todos los datos
	// dado que quizas actualice algunos
	nAux, err := contacto.GetContacto(conexion.Db)
	if err != nil {
		// Contacto no existe
		return
	} else {
		////////////////////////////
		// asignar los atributos leidos al actual para que solo se actualice lo que viene en el body
		contacto.Nombre = nAux.Nombre
		contacto.Apellido = nAux.Apellido
		contacto.Email = nAux.Email
		contacto.Telefono = nAux.Telefono
		contacto.Dni = nAux.Dni
		contacto.Monto = nAux.Monto
		contacto.Plazo = nAux.Plazo
		contacto.Mensaje = nAux.Mensaje
		contacto.Cbu = nAux.Cbu
		contacto.IdImagen = nAux.IdImagen
		contacto.FechaIngreso = nAux.FechaIngreso
		contacto.Estado = nAux.Estado
		contacto.FechaEstado = nAux.FechaEstado
		////////////////////////////
	}

	//decodifica lo que viene y lo coloca en n
	err = json.NewDecoder(r.Body).Decode(&contacto)
	if err != nil {
		panic(err)
	}
	////////////////////////////
	// agregar campos de fecha que no vienen en el body
	// contacto.FechaLastUpdate = time.Now()
	contacto.FechaEstado = time.Now().Add(time.Hour * -3)

	cant, err := contacto.UpdateContacto(conexion.Db)
	if err != nil {
		log.Println("error UpdateContacto")
		panic(err)
	} else {
		log.Println(cant) // cantidad de actualizados
	}
	contacto, _ = contacto.GetContacto(conexion.Db) // vuelvo a leer para que devuelva los datos de la BD

	j, err := json.Marshal(contacto)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusOK)                       // escribe el 200 Ok
		w.Write(j)                                         // escribe el json
	}
}
func (rest *ContactoRestFul) DeleteContacto(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteContacto")

	contacto.CleanContacto()

	contacto.Dni = TransformRequestToDni(r)

	//  BD Conexion y DELETE
	if err := conexion.Open(conexion.GetEsquemaPepeYa()); err != nil {
		log.Println("error Open")
	} else {
		defer conexion.Close()
	}
	//Delete en la BD,  primero buscar el id para luego actualizar los datos
	cant, err := contacto.DeleteContacto(conexion.Db)
	if err != nil {
		log.Println("error DeleteContacto")
		panic(err)
	} else {
		log.Println("DeleteContacto dni: ", contacto.Dni, " Se eliminaron : ", cant, " registros")
	}
	w.WriteHeader(http.StatusNoContent) // escribe el 204 Delete
}
