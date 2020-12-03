package main

import (
	"fmt"
	"log"

	"github.com/Liberdina/protobuffers/zuvap/pagopb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	pago pagopb.Pago
)

func main() {
	pay := doPago()

	jsonDemo(pay)

}
func doPago() *pagopb.Pago {
	pay := pagopb.Pago{
		IdUsuario:   32,
		Description: "Envio paquete",
		// Pay
		PaymentType:             "TC",
		CardNumber:              "1234567812345678",
		CardExpirationDateMonth: 12,
		CardExpirationDateYear:  2030,
		CardSecurityCode:        123,
		CardHolderName:          "Emanuel Ginobili",
		Currency:                "ARS",
		Amount:                  10.50,
	}
	fmt.Println(pay.String())

	return &pay
}

/*
	// Output
		google.protobuf.Timestamp datetime_process = 11; // fecha hora de procesamiento
		int32   payment_id = 12;            // el que nos vino en el CREATE
		int32   installmentPlanDetailId = 13; // Id que vino en el GET
		int32   error_code = 14;
		string  error_message = 15;
		string  id = 16;             // Indica el id con el cual se registro el pago realizado.
		string  date

*/
func jsonDemo(pm proto.Message) {
	pmAsString := toJSON(pm)
	fmt.Println(pmAsString)

	pm2 := &pagopb.Pago{}
	fromJSON(pmAsString, pm2)
	fmt.Println("Successfully created proto struct:", pm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}
