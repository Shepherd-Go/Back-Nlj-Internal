package services_test

import (
	"context"
	"testing"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	mocks "github.com/Shepherd-Go/Back-Nlj-Internal.git/mocks/db/repository"
	mocks2 "github.com/Shepherd-Go/Back-Nlj-Internal.git/mocks/utils"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/stretchr/testify/suite"
)

var (
	dataActivateEmailWhenIsCorrect = dtos.ActivateEmail{
		Password: "12345678",
	}

	dataActivateEmailWhenIsIncorrect = dtos.ActivateEmail{
		Password: "12345",
	}
)

func TestActivateEmailSuite(t *testing.T) {
	suite.Run(t, new(ActivateEmailTestSuite))
}

type ActivateEmailTestSuite struct {
	suite.Suite
	repo        *mocks.Employee
	passEncrypt *mocks2.Password
	underTest   services.ActivateEmail
}

func (suite *ActivateEmailTestSuite) SetupTest() {
	suite.repo = &mocks.Employee{}
	suite.passEncrypt = &mocks2.Password{}
	suite.underTest = services.NewActivateEmailService(suite.repo, suite.passEncrypt)
}

func (suite *ActivateEmailTestSuite) TestActivateEmail_WhenSearchEmployeeByIDFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{}, err)

	_, err := suite.underTest.ActivateEmail(ctx, dataActivateEmailWhenIsCorrect)
	suite.Error(err)

}

func (suite *ActivateEmailTestSuite) TestActivateEmail_WhenEmployeeConfirmedEmaiIsAlready() {

	isTrue := true

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{Confirmed_Email: &isTrue}, nil)

	_, err := suite.underTest.ActivateEmail(ctx, dataActivateEmailWhenIsCorrect)
	suite.Error(err)

}

func (suite *ActivateEmailTestSuite) TestActivateEmail_WhenbothPasswordsAreTheSame() {

	isFalse := false

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{Password: []byte("12345678"), Confirmed_Email: &isFalse}, nil)

	suite.passEncrypt.Mock.On("CheckPasswordHash", []byte("12345678"), dataActivateEmailWhenIsCorrect.Password).
		Return(true)

	_, err := suite.underTest.ActivateEmail(ctx, dataActivateEmailWhenIsCorrect)
	suite.Error(err)

}

func (suite *ActivateEmailTestSuite) TestActivateEmail_WhenActivateEmailFail() {

	isFalse := false

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee, Password: []byte("1234567"), Confirmed_Email: &isFalse}, nil)

	suite.passEncrypt.Mock.On("CheckPasswordHash", []byte("1234567"), dataActivateEmailWhenIsCorrect.Password).
		Return(false)

	suite.passEncrypt.Mock.On("HashPassword", &dataActivateEmailWhenIsCorrect.Password).
		Return("12345678")

	suite.repo.Mock.On("ActivateEmail", ctx, idEmployee.String(), dataActivateEmailWhenIsCorrect.Password).
		Return(err)

	_, err := suite.underTest.ActivateEmail(ctx, dataActivateEmailWhenIsCorrect)
	suite.Error(err)

}

func (suite *ActivateEmailTestSuite) TestActivateEmail_WhenSuccessfull() {

	isFalse := false

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee, Password: []byte("1234567"), Confirmed_Email: &isFalse}, nil)

	suite.passEncrypt.Mock.On("CheckPasswordHash", []byte("1234567"), dataActivateEmailWhenIsCorrect.Password).
		Return(false)

	suite.passEncrypt.Mock.On("HashPassword", &dataActivateEmailWhenIsCorrect.Password).
		Return("12345678")

	suite.repo.Mock.On("ActivateEmail", ctx, idEmployee.String(), dataActivateEmailWhenIsCorrect.Password).
		Return(nil)

	_, err := suite.underTest.ActivateEmail(ctx, dataActivateEmailWhenIsCorrect)
	suite.NoError(err)

}
