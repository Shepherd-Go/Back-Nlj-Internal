package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/controllers"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	mocks "github.com/Shepherd-Go/Back-Nlj-Internal.git/mocks/services"
	mocks2 "github.com/Shepherd-Go/Back-Nlj-Internal.git/mocks/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	WhenProccessIsSuccess = dtos.Login{
		Identifier: "neiferjr15@gmail.com",
		Password:   "12345678",
	}

	WhenDataIsIncorrect = dtos.Login{
		Identifier: "runbex13@gmail.com",
		Password:   "87654321",
	}

	WhenMissingOneData = dtos.Login{
		Identifier: "neiferjr15@gmail.com",
		Password:   "",
	}
)

func TestSessionControllerSuit(t *testing.T) {
	suite.Run(t, new(SessionControllerTestSuite))
}

type SessionControllerTestSuite struct {
	suite.Suite
	svc       *mocks.Session
	jwt       *mocks2.JWT
	underTest controllers.Session
}

func (suite *SessionControllerTestSuite) SetupTest() {
	suite.svc = &mocks.Session{}
	suite.jwt = &mocks2.JWT{}
	suite.underTest = controllers.NewSessionController(suite.svc, suite.jwt)
}

func (suite *SessionControllerTestSuite) TestWhenBindFail() {

	body, _ := json.Marshal("")
	setupCase := SetupControllerCase(http.MethodPost, "/session", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("Session", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, nil)

	suite.jwt.Mock.On("SignedLoginToken", dtos.EmployeeResponse{}).
		Return("TOKEN", nil)

	err := suite.underTest.Session(setupCase.Ctx)
	suite.Error(err)

}

func (suite *SessionControllerTestSuite) TestWhenValidateFail() {

	body, _ := json.Marshal(WhenMissingOneData)
	setupCase := SetupControllerCase(http.MethodPost, "/session", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("Session", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, nil)

	suite.jwt.Mock.On("SignedLoginToken", dtos.EmployeeResponse{}).
		Return("TOKEN", nil)

	err := suite.underTest.Session(setupCase.Ctx)
	suite.Error(err)

}

func (suite *SessionControllerTestSuite) TestWhenServerReturnError() {

	body, _ := json.Marshal(WhenProccessIsSuccess)
	setupCase := SetupControllerCase(http.MethodPost, "/session", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("Session", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, errors.New("Error"))

	suite.jwt.Mock.On("SignedLoginToken", dtos.EmployeeResponse{}).
		Return("TOKEN", nil)

	err := suite.underTest.Session(setupCase.Ctx)
	suite.Error(err)

}

func (suite *SessionControllerTestSuite) TestWhenSignedLoginTokenReturnError() {

	body, _ := json.Marshal(WhenProccessIsSuccess)
	setupCase := SetupControllerCase(http.MethodPost, "/session", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("Session", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, nil)

	suite.jwt.Mock.On("SignedLoginToken", dtos.EmployeeResponse{}).
		Return("", errors.New("Error"))

	err := suite.underTest.Session(setupCase.Ctx)
	suite.Error(err)

}

func (suite *SessionControllerTestSuite) TestWhenProcessIsSuccess() {

	body, _ := json.Marshal(WhenProccessIsSuccess)
	setupCase := SetupControllerCase(http.MethodPost, "/session", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("Session", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, nil)

	suite.jwt.Mock.On("SignedLoginToken", dtos.EmployeeResponse{}).
		Return("TOKEN", nil)

	err := suite.underTest.Session(setupCase.Ctx)
	suite.ErrorIs(err, nil)

}
