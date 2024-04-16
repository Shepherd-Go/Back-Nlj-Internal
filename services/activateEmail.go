package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/repository"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ActivateEmail interface {
	ActivateEmail(ctx context.Context, pass dtos.ActivateEmail) (dtos.EmployeeResponse, error)
}

type activateemail struct {
	repo        repository.Employee
	passEncrypt utils.Password
}

func NewActivateEmailService(repo repository.Employee, passEncrypt utils.Password) ActivateEmail {
	return &activateemail{repo, passEncrypt}
}

func (a *activateemail) ActivateEmail(ctx context.Context, pass dtos.ActivateEmail) (dtos.EmployeeResponse, error) {

	id := ctx.Value("id").(string)

	empl, err := a.repo.SearchEmployeeByID(ctx, uuid.MustParse(id))
	if err != nil {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "unexpected server error"})
	}

	if *empl.Confirmed_Email {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "this employee already has the confirmed email"})
	}

	if a.passEncrypt.CheckPasswordHash(empl.Password, pass.Password) {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "the password is the same as the temporary one"})
	}

	a.passEncrypt.HashPassword(&pass.Password)

	if err := a.repo.ActivateEmail(ctx, id, pass.Password); err != nil {
		fmt.Println(err)
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "unexpected server error"})
	}

	isTrue := true
	empl.Confirmed_Email = &isTrue
	empl.Password = nil

	return empl, nil
}
