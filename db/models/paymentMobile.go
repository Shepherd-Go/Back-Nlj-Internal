package models

import (
	"time"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/google/uuid"
)

type PaymentsMobiles []PaymentMobile

type PaymentMobile struct {
	ID          uuid.UUID
	PaymentType string
	Bank        string
	Phone       string
	PaymentCard string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	EmployeeID  uuid.UUID
}

func (pm *PaymentMobile) BuildRegisterPaymentMobile(paymobile dtos.PaymentMobile) {
	pm.ID = paymobile.ID
	pm.PaymentType = "Pago Móvil"
	pm.Bank = paymobile.Bank
	pm.Phone = paymobile.Phone
	pm.PaymentCard = paymobile.Payment_Card
	pm.EmployeeID = paymobile.EmployeeID
}

func (pm *PaymentMobile) ToDomainDTO() dtos.PaymentMobileResponse {
	return dtos.PaymentMobileResponse{
		ID:           pm.ID,
		Payment_Type: pm.PaymentType,
		Bank:         pm.Bank,
		Phone:        pm.Phone,
		Payment_Card: pm.PaymentCard,
	}
}

func (pms *PaymentsMobiles) ToDomainDTO() dtos.PaymentsMobiles {
	var paymentMobile dtos.PaymentsMobiles

	for _, v := range *pms {
		paymentMobile.Add(v.ToDomainDTO())
	}

	return paymentMobile
}
