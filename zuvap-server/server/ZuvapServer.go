package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Liberdina/protobuffers/zuvap/pagopb"
	"github.com/Liberdina/zuvap-server/services"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	pago       pagopb.Pago
	callerRest services.CallerRest

	serverPago *grpc.Server
)

func init() {
	fmt.Println("Starting Zuvap Server")
	fmt.Println("Loading Environment Variables")
	gotenv.Load()
	callerRest.Inicializar()
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
	fmt.Println("Connecting to Database")
	//	conexion.ConectToDB()

	fmt.Println("Opening PAGOS @TCP:50051")
	lis, err := net.Listen("tcp", callerRest.GetIP())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serverPago = grpc.NewServer()
	pagopb.RegisterPagoServiceServer(serverPago, &server{})

	reflection.Register(serverPago)

	fmt.Printf("Listo para procesar en %v", callerRest.GetIP())

	if err := serverPago.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type server struct{}

func (*server) Pay(ctx context.Context, req *pagopb.PagoRequest) (*pagopb.PagoResponse, error) {
	//	log.Println("llego peticion de Pago " + req.GetPago().String())

	if req.GetPago().Description == "StopZuvapServer" {
		serverPago.Stop()
		return nil, nil
	} else {
		return realizarPago(req)
	}
}
func realizarPago(req *pagopb.PagoRequest) (*pagopb.PagoResponse, error) {
	pagoResponse, err := callerRest.DoCaller(req)
	return pagoResponse, err
}
