package repository

import (
	"context"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"gorm.io/gorm"
)

type LogIn interface {
	SearchEmployeByEmailOrUsername(ctx context.Context, identifier string) (dtos.Session, error)
}

type logIn struct {
	db *gorm.DB
}

func NewLogInRepository(db *gorm.DB) LogIn {
	return &logIn{db}
}

func (e *logIn) SearchEmployeByEmailOrUsername(ctx context.Context, identifier string) (dtos.Session, error) {

	empl := models.Session{}

	if err := e.db.WithContext(ctx).Table("employees").
		Where("username=?", identifier).Or("email=?", identifier).Not("deleted=?", true).
		Select("id, first_name, last_name, password, permissions, confirmed_email, status").
		Scan(&empl).Error; err != nil {
		return dtos.Session{}, err
	}

	return empl.ToDomainDTO(), nil

}
