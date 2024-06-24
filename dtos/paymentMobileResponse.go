package dtos

import (
	"github.com/google/uuid"
)

type PaymentsMobiles []PaymentMobileResponse

type PaymentMobileResponse struct {
	ID           uuid.UUID `json:"id"`
	Payment_Type string    `json:"payment_type"`
	Bank         string    `json:"bank"`
	Phone        string    `json:"phone"`
	Payment_Card string    `json:"payment_card"`
}

func (pms *PaymentsMobiles) Add(paymentMobile PaymentMobileResponse) {
	*pms = append(*pms, paymentMobile)
}
