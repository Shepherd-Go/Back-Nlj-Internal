package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	mocks "github.com/Shepherd-Go/Back-Nlj-Internal.git/mocks/db/repository"
	mocks2 "github.com/Shepherd-Go/Back-Nlj-Internal.git/mocks/utils"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

const idKey, permissionsKey = "id", "permissions"

var (
	ctx = context.Background()
	err = errors.New("Error")

	isTrue = true

	idEmployee = uuid.MustParse("37a35c67-4953-468f-8d9a-75d4bd0c673b")

	dataCreateEmployeeIsCorrect = dtos.RegisterEmployee{
		ID:           idEmployee,
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@test.net",
		Phone:        "00000000000",
		Password:     "testtest",
		Permissions:  "administrator",
		Code_Bank:    "test",
		Pay_Phone:    "00000000000",
		Payment_Card: "00000000",
		Created_by:   "37a35c67-4953-468f-8d9a-75d4bd0c673b",
		Updated_by:   "37a35c67-4953-468f-8d9a-75d4bd0c673b",
	}

	dataUpdateEmployeeIsCorrect = dtos.UpdateEmployee{
		ID:           idEmployee,
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@test.com",
		Phone:        "00000000000",
		Permissions:  "seller",
		Code_Bank:    "0000 (test)",
		Pay_Phone:    "00000000000",
		Payment_Card: "v0000000",
		Status:       &isTrue,
	}

	dataUpdateEmployeeIsIncorrect = dtos.UpdateEmployee{
		ID:           idEmployee,
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@test.com",
		Phone:        "00000000000",
		Permissions:  "error",
		Code_Bank:    "0000 (test)",
		Pay_Phone:    "00000000000",
		Payment_Card: "v0000000",
		Status:       &isTrue,
	}
)

func TestEmployeeServiceSuite(t *testing.T) {
	suite.Run(t, new(EmployeeServiceTestSuite))
}

type EmployeeServiceTestSuite struct {
	suite.Suite
	repo      *mocks.Employee
	pass      *mocks2.Password
	logs      *mocks2.LogsError
	underTest services.Employee
}

func (suite *EmployeeServiceTestSuite) SetupTest() {
	suite.repo = &mocks.Employee{}
	suite.pass = &mocks2.Password{}
	suite.logs = &mocks2.LogsError{}
	suite.underTest = services.NewServiceEmployee(suite.repo, suite.pass, suite.logs)
}

func (suite *EmployeeServiceTestSuite) TestCreate_WhenSearchEmployeeByEmailFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByEmail", ctx, dataCreateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, err)

	dataCreateEmployeeIsCorrect.Permissions = "test"

	suite.Error(suite.underTest.CreateEmployee(ctx, dataCreateEmployeeIsCorrect))

}

func (suite *EmployeeServiceTestSuite) TestCreate_WhenEmployeeExists() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByEmail", ctx, dataCreateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.Error(suite.underTest.CreateEmployee(ctx, dataCreateEmployeeIsCorrect))

}

func (suite *EmployeeServiceTestSuite) TestCreate_WhenParsePermissionsFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByEmail", ctx, dataCreateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	suite.pass.Mock.On("GenerateTemporaryPassword").
		Return("testtest")
	suite.pass.Mock.On("HashPassword", &dataCreateEmployeeIsCorrect.Password)

	dataCreateEmployeeIsCorrect.Permissions = "admin"

	suite.Error(suite.underTest.CreateEmployee(ctx, dataCreateEmployeeIsCorrect))

}

func (suite *EmployeeServiceTestSuite) TestCreate_WhenCreateEmployeeFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByEmail", ctx, dataCreateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	suite.pass.Mock.On("GenerateTemporaryPassword").
		Return("testtest")
	suite.pass.Mock.On("HashPassword", &dataCreateEmployeeIsCorrect.Password)

	buildEmployee := models.Employee{}
	buildEmployee.BuildCreateEmployeeModel(dataCreateEmployeeIsCorrect)

	buildEmployee.Permissions = "1"

	suite.repo.Mock.On("CreateEmployee", ctx, buildEmployee).
		Return(err)

	suite.Error(suite.underTest.CreateEmployee(ctx, dataCreateEmployeeIsCorrect))

}

func (suite *EmployeeServiceTestSuite) TestCreate_WhenCreateEmployeeSuccess() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByEmail", ctx, dataCreateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	suite.pass.Mock.On("GenerateTemporaryPassword").
		Return("testtest")
	suite.pass.Mock.On("HashPassword", &dataCreateEmployeeIsCorrect.Password)

	buildEmployee := models.Employee{}
	buildEmployee.BuildCreateEmployeeModel(dataCreateEmployeeIsCorrect)

	buildEmployee.Permissions = "1"

	suite.repo.Mock.On("CreateEmployee", ctx, buildEmployee).
		Return(nil)

	suite.NoError(suite.underTest.CreateEmployee(ctx, dataCreateEmployeeIsCorrect))

}

func (suite *EmployeeServiceTestSuite) TestGet_WhenSearchAllEmployeesFail() {

	suite.repo.Mock.On("SearchAllEmployees", ctx).
		Return(dtos.Employees{}, err)

	_, err := suite.underTest.GetEmployees(ctx)
	suite.Error(err)

}

func (suite *EmployeeServiceTestSuite) TestGet_WhenSuccessfull() {

	suite.repo.Mock.On("SearchAllEmployees", ctx).
		Return(dtos.Employees{}, nil)

	_, err := suite.underTest.GetEmployees(ctx)
	suite.NoError(err)

}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenSearchEmployeeByIDFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{}, err)

	suite.Error(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsCorrect))
}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenSearchEmployeeNotExists() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	suite.Error(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsCorrect))
}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenSearchEmployeeByEmailAndNotIDFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("SearchEmployeeByEmailAndNotID", ctx, idEmployee, dataUpdateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{}, err)

	suite.Error(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsCorrect))
}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenEmailExists() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("SearchEmployeeByEmailAndNotID", ctx, idEmployee, dataUpdateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.New()}, nil)

	suite.Error(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsCorrect))
}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenParsePermissionsFail() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("SearchEmployeeByEmailAndNotID", ctx, idEmployee, dataUpdateEmployeeIsIncorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	suite.Error(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsIncorrect))
}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenUpdateEmployeeFail() {

	idToken := uuid.New()

	ctx = context.WithValue(ctx, idKey, idToken.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("SearchEmployeeByEmailAndNotID", ctx, idEmployee, dataUpdateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	dataUpdateEmployeeIsCorrect.ID = idEmployee
	dataUpdateEmployeeIsCorrect.Updated_by = idToken.String()

	buildModelEmploye := models.Employee{}
	buildModelEmploye.BuildUpdatedEmployeeModel(dataUpdateEmployeeIsCorrect)

	buildModelEmploye.Permissions = "2"

	suite.repo.Mock.On("UpdateEmployee", ctx, buildModelEmploye, idEmployee).
		Return(err)

	suite.Error(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsCorrect))
}

func (suite *EmployeeServiceTestSuite) TestUpdate_WhenUpdateEmployeeSuccess() {

	idToken := uuid.New()

	ctx = context.WithValue(ctx, idKey, idToken.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("SearchEmployeeByEmailAndNotID", ctx, idEmployee, dataUpdateEmployeeIsCorrect.Email).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	dataUpdateEmployeeIsCorrect.ID = idEmployee
	dataUpdateEmployeeIsCorrect.Updated_by = idToken.String()

	buildModelEmploye := models.Employee{}
	buildModelEmploye.BuildUpdatedEmployeeModel(dataUpdateEmployeeIsCorrect)

	buildModelEmploye.Permissions = "2"

	suite.repo.Mock.On("UpdateEmployee", ctx, buildModelEmploye, idEmployee).
		Return(nil)

	suite.NoError(suite.underTest.UpdateEmployees(ctx, idEmployee, dataUpdateEmployeeIsCorrect))
}

func (suite *EmployeeServiceTestSuite) TestDelete_WhenEmployeeIsTheEmployeeDelete() {

	ctx = context.WithValue(ctx, idKey, idEmployee.String())

	suite.Error(suite.underTest.DeleteEmployee(ctx, idEmployee))

}

func (suite *EmployeeServiceTestSuite) TestDelete_WhenSearchEmployeeByIDFail() {

	ctx = context.WithValue(ctx, idKey, uuid.New().String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, err)

	suite.Error(suite.underTest.DeleteEmployee(ctx, idEmployee))

}

func (suite *EmployeeServiceTestSuite) TestDelete_WhenEmployeeNotExists() {

	ctx = context.WithValue(ctx, idKey, uuid.New().String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: uuid.Nil}, nil)

	suite.Error(suite.underTest.DeleteEmployee(ctx, idEmployee))

}

func (suite *EmployeeServiceTestSuite) TestDelete_WhenDeleteEmployeeFail() {

	idToken := uuid.New()

	ctx = context.WithValue(ctx, idKey, idToken.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("DeleteEmployee", ctx, idEmployee, idToken.String()).
		Return(err)

	suite.Error(suite.underTest.DeleteEmployee(ctx, idEmployee))

}

func (suite *EmployeeServiceTestSuite) TestDelete_WhenDeleteEmployeeSuccess() {

	idToken := uuid.New()

	ctx = context.WithValue(ctx, idKey, idToken.String())

	suite.repo.Mock.On("SearchEmployeeByID", ctx, idEmployee).
		Return(dtos.EmployeeResponse{ID: idEmployee}, nil)

	suite.repo.Mock.On("DeleteEmployee", ctx, idEmployee, idToken.String()).
		Return(nil)

	suite.NoError(suite.underTest.DeleteEmployee(ctx, idEmployee))

}
