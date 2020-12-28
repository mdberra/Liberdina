package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Liberdina/protobuffers/zuvap/pagopb"
	"github.com/Liberdina/zuvap-server/connect"
	"github.com/Liberdina/zuvap-server/param"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	importeStr        string
	expiracionMesStr  string
	expiracionAnioStr string
	securityCodeStr   string
)

type PagosService struct {
	param    param.Parameters
	conexion connect.Conexion
}

func (c *PagosService) SetParameters(param *param.Parameters, conexion *connect.Conexion) {
	c.param = *param
	c.conexion = *conexion
}

func (c *PagosService) Pagar(w http.ResponseWriter, r *http.Request) {
	log.Println("Hice un POST")
	// antes de mover los datos blanqueo la salida porque sino queda de la vez anterior
	pago.CleanPago()

	//decodifica lo que viene
	err := json.NewDecoder(r.Body).Decode(&pago)
	if err != nil {
		panic(err)
	}
	//-------------------------------------------------------------
	req := &pagopb.PagoRequest{
		Pago: &pagopb.Pago{
			IdUsuario:   pago.IdUsuario,
			Description: pago.Description,
			// Pay
			PaymentType:             pago.PaymentType,
			CardNumber:              pago.CardNumber,
			CardExpirationDateMonth: pago.CardExpirationDateMonth,
			CardExpirationDateYear:  pago.CardExpirationDateYear,
			CardSecurityCode:        pago.CardSecurityCode,
			CardHolderName:          pago.CardHolderName,
			Currency:                pago.Currency,
			Amount:                  pago.Amount,
		},
	}
	//---------------------------------------------------------------
	idPago, err := c.conexion.LoguearInsert(req)
	if err != nil {
		fmt.Printf("NO se esta logueando en DB INSERT: %v, ERROR %v", req, err)
	}
	pagoResponse, err := c.DoCaller(req)

	err = c.conexion.LoguearUpdate(idPago, pagoResponse)
	if err != nil {
		fmt.Printf("NO se esta logueando en DB UPDATE: IdPago: %v, ERROR: %v", idPago, err)
	}
	//-------------------------------------------------------------
	pago.DatetimeProcess = pagoResponse.Pago.DatetimeProcess.AsTime()
	pago.PaymentId = pagoResponse.Pago.PaymentId
	pago.InstallmentPlanDetailId = int(pagoResponse.Pago.InstallmentPlanDetailId)
	pago.ErrorCode = int(pagoResponse.Pago.ErrorCode)
	pago.ErrorMessage = pagoResponse.Pago.ErrorMessage

	j, err := json.Marshal(pago)
	if err != nil {
		log.Println("problema en el json.Marshal")
	} else {
		w.Header().Set("Content-Type", "application/json") // standard del protocolo http
		w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
		w.Write(j)                                         // escribe el json
	}
}
func (c *PagosService) DoCaller(req *pagopb.PagoRequest) (*pagopb.PagoResponse, error) {
	pagoResponse.Pago = req.GetPago()
	importeStr = fmt.Sprintf("%.2f", req.GetPago().GetAmount())
	expiracionMesStr = fmt.Sprintf("%2d", req.GetPago().GetCardExpirationDateMonth())
	expiracionAnioStr = fmt.Sprintf("%2d", req.GetPago().GetCardExpirationDateYear())
	securityCodeStr = fmt.Sprintf("%3d", req.GetPago().GetCardSecurityCode())

	if err := c.doCreate(req); err != nil {
		return &pagoResponse, err
	}
	if err := c.doPlanDetail(req); err != nil {
		return &pagoResponse, err
	}
	if err := c.doPay(req); err != nil {
		return &pagoResponse, err
	}

	return &pagoResponse, nil
}

type CreateBody struct {
	SuccessUrl    string            `json:"success_url"`
	FailureUrl    string            `json:"failure_url"`
	CurrencyCode  string            `json:"currency_code"`
	CustomerEmail string            `json:"customer_email"`
	Description   string            `json:"description"`
	ItemList      [1]CreateItemList `json:"item_list"`
	TotalAmount   float64           `json:"total_amount"`
	Token         string            `json:"token"`
	OperationId   string            `json:"operation_id"`
}
type CreateItemList struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Quantity    float64 `json:"quantity"`
}
type CreateRetorno struct {
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	Url            string `json:"url"`
	ExpirationDate string `json:"expiration_date"`
	PaymentId      string `json:"payment_id"`
}

func (c *PagosService) doCreate(req *pagopb.PagoRequest) error {
	timeNowStr := time.Now().Format(time.RFC3339)
	usuarioStr := fmt.Sprint(req.GetPago().GetIdUsuario()) + "@" + timeNowStr

	var cil [1]CreateItemList
	cil[0].Description = req.GetPago().GetDescription()
	cil[0].Amount = req.GetPago().GetAmount()
	cil[0].Quantity = 1
	var cb CreateBody
	cb.SuccessUrl = ""
	cb.FailureUrl = ""
	cb.CurrencyCode = req.GetPago().GetCurrency()
	cb.CustomerEmail = ""
	cb.Description = req.GetPago().GetDescription()
	cb.ItemList = cil
	cb.TotalAmount = req.GetPago().GetAmount()
	cb.Token = c.param.Token
	cb.OperationId = usuarioStr

	requestBody, err := json.Marshal(cb)
	if err != nil {
		//		log.Printf("Error al hacer son Marshal en create cb %v ", cb)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer son Marshal en create cb %v ", cb),
		)
	}
	//	fmt.Println(string(requestBody))
	//--------------------------------------------------------------------------
	create, err := c.DoRest("POST", c.param.URLPostCreate, requestBody)
	if err != nil {
		//		log.Printf("Error al hacer Post Create %v", requestBody)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer Post Create %v", requestBody),
		)
	}
	//	fmt.Println(string(create))

	var cr CreateRetorno
	err = json.Unmarshal(create, &cr)
	if err != nil {
		//		log.Printf("Error al hacer Unmarshal en Post create %v  err %v", create, err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer Unmarshal en Post create %v  err %v", create, err),
		)
	}
	//	fmt.Println(cr)

	if !cr.Success {
		//		return errors.New(fmt.Sprintf("DoCreate: success: %v message: %v", cr.Success, cr.Message))
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("DoCreate: success: %v message: %v", cr.Success, cr.Message),
		)
	}
	//	fmt.Println(cr.PaymentId)
	pagoResponse.Pago.PaymentId = cr.PaymentId
	if cr.Success {
		pagoResponse.Pago.ErrorCode = 200
	} else {
		pagoResponse.Pago.ErrorCode = 0
	}
	pagoResponse.Pago.ErrorMessage = cr.Message
	return nil
}

//-------------------------------------------------------------------------------
type PlanDetailsRetorno struct {
	ErrorCode         int32    `json:"error_code"`
	ErrorMessages     []string `json:"error_messages"`
	Id                int32    `json:"id"`
	Quantity          int32    `json:"quantity"`
	TotalAmount       float64  `json:"total_amount"`
	InstallmentAmount float64  `json:"installment_amount"`
}

func (c *PagosService) doPlanDetail(req *pagopb.PagoRequest) error {
	//	Parametros requeridos por query string:
	//	[string] cardNumber: Primeros 6 números de la tarjeta de crédito.
	//	[string] amount: Monto neto de la operación
	//	?cardNumber=427602&amount=1000
	result := "?cardNumber=" + req.GetPago().GetCardNumber()[0:6] + "&amount=" + importeStr

	httpget, err := c.DoRest("GET", c.param.URLGetPlanDetails+result, nil)
	if err != nil {
		//		log.Printf("Error al hacer Get PlanDetail %v", result)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer Get PlanDetail %v", result),
		)
	}
	//	fmt.Println(string(httpget))

	var pdr []PlanDetailsRetorno
	err = json.Unmarshal(httpget, &pdr)
	if err != nil {
		log.Printf("Error al hacer son Unmarshal en Get pdr %v  err %v", pdr, err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer son Unmarshal en Get pdr %v  err %v", pdr, err),
		)
	}
	//	fmt.Println(pdr)
	if pdr[0].ErrorCode != 0 && pdr[0].ErrorCode != 200 {
		pagoResponse.Pago.ErrorCode = pdr[0].ErrorCode
		pagoResponse.Pago.ErrorMessage = pdr[0].ErrorMessages[0]
		//		return errors.New(fmt.Sprintf("DoPlanDetail: error_code: %v error_message: %v", pdr[0].ErrorCode, pdr[0].ErrorMessages))
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf(fmt.Sprintf("DoPlanDetail: error_code: %v error_message: %v", pdr[0].ErrorCode, pdr[0].ErrorMessages)),
		)
	}
	//	fmt.Println(pdr[0].Id)
	pagoResponse.Pago.InstallmentPlanDetailId = pdr[0].Id
	return nil
}

//-------------------------------------------------------------------------------
type PayBody struct {
	Currency                 string                   `json:"currency"`
	Amount                   float64                  `json:"amount"`
	CardAmount               float64                  `json:"cardAmount"`
	PaymentId                string                   `json:"paymentId"`
	PaymentMethodDescription string                   `json:"paymentMethodDescription"`
	InstallmentPlanDetailId  int32                    `json:"installmentPlanDetailId"` //Id del plan de cuotas obtenido con la consulta de Planes de cuotas disponibles.
	SourcePayments           [1]PaySourcePaymentsList `json:"sourcePayments"`          //Información sobre los medio de pago a utilizar.
}
type PaySourcePaymentsList struct {
	PaymentType             string  `json:"paymentType"`             // Tipo de medio de pago a utilizar. (TC / TD)
	CardNumber              string  `json:"cardNumber"`              // Número de tarjeta de crédito. Sin espacios.
	CardExpirationDateMonth string  `json:"cardExpirationDateMonth"` // Mes de expiración de la tarjeta a utilizar.
	CardExpirationDateYear  string  `json:"cardExpirationDateYear"`  // Año de expiración de la tarjeta a utilizar.
	CardSecurityCode        string  `json:"cardSecurityCode"`        // Código de seguridad de la tarjeta a utilizar
	CardHolderName          string  `json:"cardHolderName"`          // Nombre completo que figura en la tarjeta a utilizar.
	Amount                  float64 `json:"amount"`                  // Importe bruto de la operación, es decir, neto + costo del plan de cuotas seleccionado.
}
type PayRetorno struct {
	ErrorCode     int32    `json:"error_code"`
	ErrorMessages []string `json:"error_messages"`
	Id            int32    `json:"id"`   // Indica el id con el cual se registro el pago realizado.
	Date          string   `json:"date"` // Indica la fecha en la cual se registro el pago realizado.
}

func (c *PagosService) doPay(req *pagopb.PagoRequest) error {
	var pspl [1]PaySourcePaymentsList
	pspl[0].PaymentType = req.GetPago().GetPaymentType()
	pspl[0].CardNumber = req.GetPago().GetCardNumber()
	pspl[0].CardExpirationDateMonth = expiracionMesStr
	pspl[0].CardExpirationDateYear = expiracionAnioStr
	pspl[0].CardSecurityCode = securityCodeStr
	pspl[0].CardHolderName = req.GetPago().GetCardHolderName()
	pspl[0].Amount = req.GetPago().GetAmount()
	var pb PayBody
	pb.Currency = req.GetPago().GetCurrency()
	pb.Amount = req.GetPago().GetAmount()
	pb.CardAmount = req.GetPago().GetAmount()
	pb.PaymentId = pagoResponse.Pago.PaymentId
	pb.PaymentMethodDescription = req.GetPago().GetPaymentType() + req.GetPago().GetCardNumber()[0:4]
	pb.InstallmentPlanDetailId = pagoResponse.Pago.InstallmentPlanDetailId
	pb.SourcePayments = pspl

	requestBody, err := json.Marshal(pb)
	//	fmt.Println(string(requestBody))
	if err != nil {
		//		log.Printf("Error al hacer son Marshal en pay cb %v ", pb)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer son Marshal en pay cb %v ", pb),
		)
	}
	//	fmt.Println(string(requestBody))
	//--------------------------------------------------------------------------
	pay, err := c.DoRest("POST", c.param.URLPostPay, requestBody)
	if err != nil {
		//		log.Printf("Error al hacer Post Pay %v", requestBody)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer Post Pay %v", requestBody),
		)
	}
	//	fmt.Println(string(pay))

	var pr PayRetorno
	err = json.Unmarshal(pay, &pr)
	if err != nil {
		//		log.Printf("Error al hacer son Unmarshal en Post Pay %v  err %v", pay, err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error al hacer son Unmarshal en Post Pay %v  err %v", pay, err),
		)
	}
	//	fmt.Println(pr)
	if pr.ErrorCode != 0 && pr.ErrorCode != 200 {
		pagoResponse.Pago.ErrorCode = pr.ErrorCode
		pagoResponse.Pago.ErrorMessage = pr.ErrorMessages[0]
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("DoPay: error_code: %v error_message: %v", pr.ErrorCode, pr.ErrorMessages),
		)
		//		return errors.New(fmt.Sprintf("DoPay: error_code: %v error_message: %v", pr.ErrorCode, pr.ErrorMessages))
	}
	pagoResponse.Pago.PagoId = fmt.Sprint(pr.Id)
	pagoResponse.Pago.PagoDate = pr.Date
	return nil
}

//-----------------------------------------------------------------------------------------
func (*PagosService) DoRest(method string, url string, requestBody []byte) ([]byte, error) {
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Printf("Error al hacer el http.NewRequest %v", err)
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Printf("Error al hacer el Request %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al hacer el ioutil.ReadAll %v", err)
		return nil, err
	}
	return body, nil
}
