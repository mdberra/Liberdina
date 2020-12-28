package data

import (
	"time"
)

type Pago struct {
	IdUsuario   int32  `json:"idUsuario"`
	Description string `json:"description"`
	// Pay
	PaymentType             string  `json:"paymentType"`
	CardNumber              string  `json:"cardNumber"`
	CardExpirationDateMonth int32   `json:"cardExpirationDateMonth"`
	CardExpirationDateYear  int32   `json:"cardExpirationDateYear"`
	CardSecurityCode        int32   `json:"cardSecurityCode"`
	CardHolderName          string  `json:"cardHolderName"`
	Currency                string  `json:"currency"`
	Amount                  float64 `json:"amount"`
	// Output
	DatetimeProcess         time.Time `json:"datetimeProcess"`
	PaymentId               string    `json:"paymentId"`
	InstallmentPlanDetailId int       `json:"installmentPlanDetailId"`
	ErrorCode               int       `json:"errorCode"`
	ErrorMessage            string    `json:"errorMessage"`
}

func (p *Pago) CleanPago() {
	p.IdUsuario = 0
	p.Description = ""
	p.PaymentType = ""
	p.CardNumber = ""
	p.CardExpirationDateMonth = 0
	p.CardExpirationDateYear = 0
	p.CardSecurityCode = 0
	p.CardHolderName = ""
	p.Currency = ""
	p.Amount = 0
	p.DatetimeProcess = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	p.PaymentId = ""
	p.InstallmentPlanDetailId = 0
	p.ErrorCode = 0
	p.ErrorMessage = ""
}
