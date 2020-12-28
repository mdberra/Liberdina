package services

import (
	"github.com/Liberdina/protobuffers/zuvap/pagopb"
	"github.com/Liberdina/zuvap-server/connect"
	"github.com/Liberdina/zuvap-server/data"
)

var (
	pago         data.Pago
	conexion     connect.Conexion
	pagoResponse pagopb.PagoResponse
)
