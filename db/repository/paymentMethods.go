package repository

import (
	"context"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethods interface {
	RegisterPaymentMobile(ctx context.Context, payMobl models.PaymentMobile) error
	SearchPaymentMobileByID(ctx context.Context, id uuid.UUID) (dtos.PaymentMobileResponse, error)
	SearchPaymentMobileByEmployeeID(ctx context.Context, idEmployee uuid.UUID) (dtos.PaymentMobileResponse, error)
}

type paymentmethods struct {
	db *gorm.DB
}

func NewPaymentMethodsRepository(db *gorm.DB) PaymentMethods {
	return &paymentmethods{db}
}

func (pmthds *paymentmethods) RegisterPaymentMobile(ctx context.Context, payMobl models.PaymentMobile) error {

	if err := pmthds.db.WithContext(ctx).Table("payment_mobile").Create(&payMobl).Error; err != nil {
		return err
	}

	return nil
}

func (pmthds *paymentmethods) SearchPaymentMobileByID(ctx context.Context, id uuid.UUID) (dtos.PaymentMobileResponse, error) {

	paymobl := models.PaymentMobile{}

	if err := pmthds.db.WithContext(ctx).Table("payment_mobile").Where("id=?", id).
		Select("id, payment_type", "bank", "phone", "payment_card").
		Scan(&paymobl).
		Error; err != nil {
		return dtos.PaymentMobileResponse{}, err
	}

	return paymobl.ToDomainDTO(), nil
}

func (pmthds *paymentmethods) SearchPaymentMobileByEmployeeID(ctx context.Context, idEmployee uuid.UUID) (dtos.PaymentMobileResponse, error) {

	paymobl := models.PaymentMobile{}

	if err := pmthds.db.WithContext(ctx).Table("payment_mobile").Where("employee_id=?", idEmployee).
		Select("id, payment_type", "bank", "phone", "payment_card").
		Scan(&paymobl).
		Error; err != nil {
		return dtos.PaymentMobileResponse{}, err
	}

	return paymobl.ToDomainDTO(), nil
}
