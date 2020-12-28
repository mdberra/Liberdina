package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Liberdina/protobuffers/zuvap/pagopb"
	"github.com/Liberdina/zuvap-server/connect"
	"github.com/Liberdina/zuvap-server/server/param"
	"github.com/Liberdina/zuvap-server/services"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	pago       pagopb.Pago
	callerRest services.CallerRest
	parameters param.Parameters
	serverPago *grpc.Server
	conexion   connect.Conexion
)

func init() {
	fmt.Println("Starting Zuvap Server")
	fmt.Println("Loading Environment Variables")
	gotenv.Load("parametros.env")
}

/*
cargar variables de entorno
hacer protobuf de entrada y retorno
recibir
GoRoutine
POST /api/paymentgateway/payment/create
GET  /api/paymentgateway/installment_plan_details
POST /api/paymentgateway/payment/pay
call StoreProcedure a PostgreSQL
Retorno
*/
func main() {
	parameters, err := parameters.SetParameters()
	if err != nil {
		log.Fatalf("Error SetParameters: %v", err)
	}

	fmt.Println("Connecting to Database")
	if err := conexion.ConectToDB(parameters); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	/*  PROTOCOL BUFFERS */
	fmt.Println("Opening PAGOS")
	lis, err := net.Listen("tcp", parameters.IP+parameters.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	callerRest.SetParameters(parameters)
	serverPago = grpc.NewServer()
	pagopb.RegisterPagoServiceServer(serverPago, &server{})

	reflection.Register(serverPago)

	fmt.Printf("Listo para procesar en %v", parameters.IP+parameters.Port)
	fmt.Println()

	if err := serverPago.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type server struct{}

func (s *server) Pay(ctx context.Context, req *pagopb.PagoRequest) (*pagopb.PagoResponse, error) {
	//	log.Println("llego peticion de Pago " + req.GetPago().String())

	fmt.Println(req)
	return s.realizarPago(req)
}
func (s *server) realizarPago(req *pagopb.PagoRequest) (*pagopb.PagoResponse, error) {
	idPago, err := conexion.LoguearInsert(req)
	if err != nil {
		fmt.Printf("NO se esta logueando en DB INSERT: %v, ERROR %v", req, err)
	}

	pagoResponse, err := callerRest.DoCaller(req)

	err = conexion.LoguearUpdate(idPago, pagoResponse)
	if err != nil {
		fmt.Printf("NO se esta logueando en DB UPDATE: IdPago: %v, ERROR: %v", idPago, err)
	}
	return pagoResponse, err
}
