package connect

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ambiente     string
	tokenTesting string = "MIIII-ca84ad9774ab42f1bd0c35d2cbee8820"
	tokenProd    string = "MIIII-a26015b135914a159427865f4ca7c3ae"

	urlTesting    string = "https://0ynvohd8r4.execute-api.us-east-2.amazonaws.com/Staging"
	urlProduccion string = "https://94lsuekt0k.execute-api.us-east-2.amazonaws.com/prod"

	postCreate     string = "/api/paymentgateway/payment/create"
	getPlanDetails string = "/api/paymentgateway/installment_plan_details"
	postPay        string = "/api/paymentgateway/payment/pay"
)

type Rest struct {
	IP                string
	Token             string
	URLPostCreate     string
	URLGetPlanDetails string
	URLPostPay        string
}

func (c *Rest) SetearAmbiente() {
	if ambiente == "" {
		ambiente = os.Getenv("Ambiente")
		if ambiente == "" {
			log.Fatalln("No se encontro el Ambiente")
			panic("ambiente")
		}
		switch ambiente {
		case "Local":
			c.IP = "Localhost:50051"
			c.Token = tokenTesting
			c.URLPostCreate = urlTesting + postCreate
			c.URLGetPlanDetails = urlTesting + getPlanDetails
			c.URLPostPay = urlTesting + postPay
		case "Test":
			c.IP = "134.209.65.78:50051"
			c.Token = tokenTesting
			c.URLPostCreate = urlTesting + postCreate
			c.URLGetPlanDetails = urlTesting + getPlanDetails
			c.URLPostPay = urlTesting + postPay
		case "Prod":
			c.IP = "142.93.122.23:50051"
			c.Token = tokenProd
			c.URLPostCreate = urlProduccion + postCreate
			c.URLGetPlanDetails = urlProduccion + getPlanDetails
			c.URLPostPay = urlProduccion + postPay

		default:
			log.Fatalf("Ambiente mal definido: %v", ambiente)
			panic("No se encontro el Ambiente" + ambiente)
		}
	}
}

func (*Rest) DoRest(method string, url string, requestBody []byte) ([]byte, error) {
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
