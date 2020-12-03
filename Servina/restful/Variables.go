package restful

import (
	"log"
	"net/http"
	"strconv"

	data "github.com/Liberdina/Servina/data"
	dataPepeya "github.com/Liberdina/Servina/data/pepeya"
	dataServina "github.com/Liberdina/Servina/data/servina"
	"github.com/Liberdina/Servina/services"
	"github.com/gorilla/mux"
)

var (
	conexion data.Conexion
	keyValue data.KeyValue
	cliente  data.Cliente

	contacto dataPepeya.Contacto

	debito       dataServina.Debito
	general      dataServina.General
	findCliente  dataServina.FindCliente
	delegaciones dataServina.Delegaciones
	delegacion   dataServina.Delegacion
	banco        dataServina.Banco
	diasCobro    dataServina.DiasCobro

	cloudStorageService services.CloudStorageService
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
