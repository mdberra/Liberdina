package RestFul

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Liberdina/Recedig/Base"
	"github.com/Liberdina/Recedig/Data"
	"github.com/Liberdina/Recedig/Services"
	"github.com/gorilla/mux"
)

var (
	conexion    Base.Conexion
	keyValue    Base.KeyValue
	customError Base.CustomError

	medico      Data.Medico
	farmacia    Data.Farmacia
	laboratorio Data.Laboratorio
	medicamento Data.Medicamento
	receta      Data.Receta

	recetaItem  Data.RecetaItem
	recetaItems []Data.RecetaItem

	cloudStorageService Services.CloudStorageService
)

func TransformRequestToDni(r *http.Request) int64 {
	vars := mux.Vars(r)
	id := vars["dni"]
	dni, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("error TransformRequestToDni %v", r)
	}
	return dni
}
