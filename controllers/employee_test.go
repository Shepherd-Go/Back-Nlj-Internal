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
	isTrue = true

	idEmployee = uuid.MustParse("37a35c67-4953-468f-8d9a-75d4bd0c673b")

	WhenDataIsValidInCreate = dtos.RegisterEmployee{
		ID:           idEmployee,
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@gmail.com",
		Phone:        "00000000000",
		Permissions:  "administrator",
		Code_Bank:    "test",
		Pay_Phone:    "00000000000",
		Payment_Card: "test",
	}

	WhenDataIsValidInUpdate = dtos.UpdateEmployee{
		ID:           uuid.New(),
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@gmail.com",
		Phone:        "00000000000",
		Permissions:  "administrator",
		Code_Bank:    "test",
		Pay_Phone:    "00000000000",
		Payment_Card: "test",
		Status:       &isTrue,
	}

	WhenDataIsBadInUpdate = dtos.UpdateEmployee{
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@gmail.com",
		Phone:        "test",
		Permissions:  "administrator",
		Code_Bank:    "test",
		Pay_Phone:    "00000000000",
		Payment_Card: "test",
		Status:       &isTrue,
	}
)

func TestEmployeeControllerSuit(t *testing.T) {
	suite.Run(t, new(EmployeeControllerTestSuit))
}

type EmployeeControllerTestSuit struct {
	suite.Suite
	svc       *mocks.Employee
	underTest controllers.Employee
}

func (suite *EmployeeControllerTestSuit) SetupTest() {
	suite.svc = &mocks.Employee{}
	suite.underTest = controllers.NewEmployeeController(suite.svc)
}

func (suite *EmployeeControllerTestSuit) TestWhenBindFailInCreate() {

	body, _ := json.Marshal("")
	setupCase := SetupControllerCase(http.MethodPost, "/employee/create", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("CreateEmployee", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, nil)

	err := suite.underTest.CreateEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenValidateFailInCreate() {

	body, _ := json.Marshal(dtos.RegisterEmployee{})
	setupCase := SetupControllerCase(http.MethodPost, "/employee", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("CreateEmployee", setupCase.Ctx.Request().Context(), WhenProccessIsSuccess).
		Return(dtos.EmployeeResponse{}, nil)

	err := suite.underTest.CreateEmployee(setupCase.Ctx)

	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenServiceFailInCreate() {

	body, _ := json.Marshal(WhenDataIsValidInCreate)
	setupCase := SetupControllerCase(http.MethodPost, "/employee/create", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("CreateEmployee", setupCase.Ctx.Request().Context(), WhenDataIsValidInCreate).
		Return(echo.NewHTTPError(422, ""))

	err := suite.underTest.CreateEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenSuccessInCreate() {

	body, _ := json.Marshal(WhenDataIsValidInCreate)
	setupCase := SetupControllerCase(http.MethodPost, "/employee/create", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("CreateEmployee", setupCase.Ctx.Request().Context(), WhenDataIsValidInCreate).
		Return(nil)

	err := suite.underTest.CreateEmployee(setupCase.Ctx)
	suite.NoError(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenServiceFailInGet() {

	setupCase := SetupControllerCase(http.MethodGet, "/employee/all", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("GetEmployees", setupCase.Ctx.Request().Context()).
		Return(dtos.Employees{}, errors.New("Error"))

	err := suite.underTest.GetEmployees(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenSuccessInGet() {

	setupCase := SetupControllerCase(http.MethodGet, "/employee/all", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("GetEmployees", setupCase.Ctx.Request().Context()).
		Return(dtos.Employees{}, nil)

	err := suite.underTest.GetEmployees(setupCase.Ctx)
	suite.NoError(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenIdConverterFailInUpdate() {

	body, _ := json.Marshal("")
	setupCase := SetupControllerCase(http.MethodPut, "/employee/update?id=1234", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("UpdateEmployees", setupCase.Ctx.Request().Context(), nil).
		Return(dtos.EmployeeResponse{}, nil)

	err := suite.underTest.UpdateEmployee(setupCase.Ctx)
	suite.Error(err)

}
func (suite *EmployeeControllerTestSuit) TestWhenBindFailInUpdate() {

	id := uuid.New().String()

	body, _ := json.Marshal("")
	setupCase := SetupControllerCase(http.MethodPut, "/employee/update?id="+id, bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("UpdateEmployees", setupCase.Ctx.Request().Context(), nil).
		Return(dtos.EmployeeResponse{}, nil)

	err := suite.underTest.UpdateEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenValidateFailInUpdate() {

	id := uuid.New().String()

	body, _ := json.Marshal(WhenDataIsBadInUpdate)
	setupCase := SetupControllerCase(http.MethodPut, "/employee/update?id="+id, bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("UpdateEmployees", setupCase.Ctx.Request().Context(), WhenDataIsBadInUpdate).
		Return(dtos.EmployeeResponse{}, nil)

	err := suite.underTest.UpdateEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenServiceFailInUpdate() {

	id := WhenDataIsValidInUpdate.ID.String()

	body, _ := json.Marshal(WhenDataIsValidInUpdate)
	setupCase := SetupControllerCase(http.MethodPut, "/employee/update?id="+id, bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("UpdateEmployees", setupCase.Ctx.Request().Context(), WhenDataIsValidInUpdate.ID, WhenDataIsValidInUpdate).
		Return(errors.New("Error"))

	err := suite.underTest.UpdateEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenSuccessInUpdate() {

	id := WhenDataIsValidInUpdate.ID.String()

	body, _ := json.Marshal(WhenDataIsValidInUpdate)
	setupCase := SetupControllerCase(http.MethodPut, "/employee/update?id="+id, bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("UpdateEmployees", setupCase.Ctx.Request().Context(), WhenDataIsValidInUpdate.ID, WhenDataIsValidInUpdate).
		Return(nil)

	err := suite.underTest.UpdateEmployee(setupCase.Ctx)
	suite.NoError(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenIdConverterFailInDelete() {

	setupCase := SetupControllerCase(http.MethodDelete, "/employee/update?id=1234", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("DeleteEmployee", setupCase.Ctx.Request().Context(), nil).
		Return(dtos.EmployeeResponse{}, nil)

	err := suite.underTest.DeleteEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenServiceFailInDelete() {

	id := uuid.New()

	setupCase := SetupControllerCase(http.MethodDelete, "/employee/delete?id="+id.String(), nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("DeleteEmployee", setupCase.Ctx.Request().Context(), id).
		Return(errors.New("Error"))

	err := suite.underTest.DeleteEmployee(setupCase.Ctx)
	suite.Error(err)

}

func (suite *EmployeeControllerTestSuit) TestWhenSuccessInDelete() {

	id := uuid.New()

	setupCase := SetupControllerCase(http.MethodDelete, "/employee/delete?id="+id.String(), nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.svc.Mock.On("DeleteEmployee", setupCase.Ctx.Request().Context(), id).
		Return(nil)

	err := suite.underTest.DeleteEmployee(setupCase.Ctx)
	suite.NoError(err)

}
