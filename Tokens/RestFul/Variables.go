package RestFul

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Tokens/Data"
	"github.com/gorilla/mux"
)

var (
	conexion Data.Conexion
	usuario  Data.Usuario
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
