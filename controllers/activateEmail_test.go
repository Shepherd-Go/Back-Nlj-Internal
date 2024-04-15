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
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	dataActivateEmailIsCorrect = dtos.ActivateEmail{
		Password: "12345678",
	}

	dataActivateEmailIsIncorrect = dtos.ActivateEmail{
		Password: "12345",
	}
)

func TestActivateControllerSuite(t *testing.T) {
	suite.Run(t, new(ActivateEmailControllerSuite))
}

type ActivateEmailControllerSuite struct {
	suite.Suite
	svc       *mocks.ActivateEmail
	underTest controllers.ActivateEmail
}

func (suite *ActivateEmailControllerSuite) SetupTest() {
	suite.svc = &mocks.ActivateEmail{}
	suite.underTest = controllers.NewActiveEmailControler(suite.svc)
}

func (suite *ActivateEmailControllerSuite) TestActivateEmail_WhenBindFail() {

	body, _ := json.Marshal("")
	setupTest := SetupControllerCase(http.MethodPut, "/employee/activate-email", bytes.NewBuffer(body))
	setupTest.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.ActivateEmail(setupTest.Ctx))

}

func (suite *ActivateEmailControllerSuite) TestActivateEmail_WhenValidateFail() {

	body, _ := json.Marshal(dataActivateEmailIsIncorrect)
	setupTest := SetupControllerCase(http.MethodPut, "/employee/activate-email", bytes.NewBuffer(body))
	setupTest.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.ActivateEmail(setupTest.Ctx))

}

func (suite *ActivateEmailControllerSuite) TestActivateEmail_WhenActivateEmailFail() {

	body, _ := json.Marshal(dataActivateEmailIsCorrect)
	setupTest := SetupControllerCase(http.MethodPut, "/employee/activate-email", bytes.NewBuffer(body))
	setupTest.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("ActivateEmail", setupTest.Ctx.Request().Context(), dataActivateEmailIsCorrect).
		Return(dtos.EmployeeResponse{}, errors.New("error"))

	suite.Error(suite.underTest.ActivateEmail(setupTest.Ctx))

}

func (suite *ActivateEmailControllerSuite) TestActivateEmail_WhenActivateEmailSuccessfull() {

	body, _ := json.Marshal(dataActivateEmailIsCorrect)
	setupTest := SetupControllerCase(http.MethodPut, "/employee/activate-email", bytes.NewBuffer(body))
	setupTest.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("ActivateEmail", setupTest.Ctx.Request().Context(), dataActivateEmailIsCorrect).
		Return(dtos.EmployeeResponse{ID: uuid.New()}, nil)

	suite.NoError(suite.underTest.ActivateEmail(setupTest.Ctx))

}
