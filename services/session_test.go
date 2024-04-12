package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/dtos"
	mocks "github.com/BBCompanyca/Back-Nlj-Internal.git/mocks/db/repository"
	mocks2 "github.com/BBCompanyca/Back-Nlj-Internal.git/mocks/utils"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

var (
	ctxTrace    = context.Background()
	errExpected = errors.New("error")

	loginData = dtos.Login{
		Identifier: "test@test.com",
		Password:   "123456",
	}
)

func TestSessionServiceSuit(t *testing.T) {
	suite.Run(t, new(SessionServiceTestSuite))
}

type SessionServiceTestSuite struct {
	suite.Suite
	repo      *mocks.Employee
	pass      *mocks2.Password
	underTest services.Session
}

func (suite *SessionServiceTestSuite) SetupTest() {
	suite.repo = &mocks.Employee{}
	suite.pass = &mocks2.Password{}
	suite.underTest = services.NewSessionService(suite.repo, suite.pass)
}

func (suite *SessionServiceTestSuite) TestWhenSearchEmployeByEmailOrUsernameFail() {

	suite.repo.Mock.On("SearchEmployeByEmailOrUsername", ctxTrace, loginData.Identifier).
		Return(dtos.EmployeeResponse{}, errExpected)

	_, err := suite.underTest.Session(ctxTrace, loginData)
	suite.Error(err)

}

func (suite *SessionServiceTestSuite) TestWhenEmployeeNotExists() {

	suite.repo.Mock.On("SearchEmployeByEmailOrUsername", ctxTrace, loginData.Identifier).
		Return(dtos.EmployeeResponse{}, nil)

	_, err := suite.underTest.Session(ctxTrace, loginData)
	suite.Error(err)

}

func (suite *SessionServiceTestSuite) TestWhenPasswordIsIncorrect() {

	suite.repo.Mock.On("SearchEmployeByEmailOrUsername", ctxTrace, loginData.Identifier).
		Return(dtos.EmployeeResponse{ID: uuid.New(), Password: []byte("12345")}, nil)

	suite.pass.Mock.On("CheckPasswordHash", []byte("12345"), loginData.Password).
		Return(false)

	_, err := suite.underTest.Session(ctxTrace, loginData)
	suite.Error(err)

}

func (suite *SessionServiceTestSuite) TestWhenStatusIsFalse() {

	isFalse := false

	suite.repo.Mock.On("SearchEmployeByEmailOrUsername", ctxTrace, loginData.Identifier).
		Return(dtos.EmployeeResponse{ID: uuid.New(), Password: []byte("123456"), Status: &isFalse}, nil)

	suite.pass.Mock.On("CheckPasswordHash", []byte("123456"), loginData.Password).
		Return(true)

	_, err := suite.underTest.Session(ctxTrace, loginData)
	suite.Error(err)

}

func (suite *SessionServiceTestSuite) TestSession_WhenSuccess() {

	isTrue := true

	suite.repo.Mock.On("SearchEmployeByEmailOrUsername", ctxTrace, loginData.Identifier).
		Return(dtos.EmployeeResponse{ID: uuid.New(), Password: []byte("123456"), Status: &isTrue}, nil)

	suite.pass.Mock.On("CheckPasswordHash", []byte("123456"), loginData.Password).
		Return(true)

	_, err := suite.underTest.Session(ctxTrace, loginData)
	suite.NoError(err)

}
