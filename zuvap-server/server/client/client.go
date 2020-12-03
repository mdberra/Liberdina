package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Liberdina/protobuffers/zuvap/pagopb"

	"google.golang.org/grpc"
)

var (
	pago pagopb.Pago
)

func main() {

	fmt.Println("Client ZuvapPagos")

	cc, err := grpc.Dial("134.209.65.78:50051", grpc.WithInsecure())
	//cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := pagopb.NewPagoServiceClient(cc)

	doUnary(c)
}

func doUnary(c pagopb.PagoServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")

	pay := pagopb.Pago{
		IdUsuario:   32,
		Description: "Compra de productos varios",
		//Description: "StopZuvapServer",
		// Pay
		PaymentType:             "TC",
		CardNumber:              "4276027812345678",
		CardExpirationDateMonth: 12,
		CardExpirationDateYear:  30,
		CardSecurityCode:        123,
		CardHolderName:          "Emanuel Ginobili",
		Currency:                "ARS",
		Amount:                  400,
	}

	req := &pagopb.PagoRequest{
		Pago: &pagopb.Pago{
			IdUsuario:   pay.IdUsuario,
			Description: pay.Description,
			// Pay
			PaymentType:             pay.PaymentType,
			CardNumber:              pay.CardNumber,
			CardExpirationDateMonth: pay.CardExpirationDateMonth,
			CardExpirationDateYear:  pay.CardExpirationDateYear,
			CardSecurityCode:        pay.CardSecurityCode,
			CardHolderName:          pay.CardHolderName,
			Currency:                pay.Currency,
			Amount:                  pay.Amount,
		},
	}

	res, err := c.Pay(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling RPC pago.Pay: %v", err)
	}
	log.Printf("Response from pago.Pay: %v", res.Pago)
}
