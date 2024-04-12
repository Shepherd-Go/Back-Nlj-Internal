package controllers_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var healthJson = `{"message":"service running successfully..."}`

type ControllerCase struct {
	Req *http.Request
	Res *httptest.ResponseRecorder
	Ctx echo.Context
}

func TestHealthControllerSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerTestSuite))
}

type HealthControllerTestSuite struct {
	suite.Suite
	underTest controllers.Health
}

func (suite *HealthControllerTestSuite) SetupTest() {
	suite.underTest = controllers.NewHealthController()
}

func SetupControllerCase(method, url string, body io.Reader) ControllerCase {
	path := fmt.Sprintf("/api%s", url)

	e := echo.New()
	req := httptest.NewRequest(method, path, body)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)

	return ControllerCase{
		req, res, ctx,
	}

}

func (suite *HealthControllerTestSuite) TestWhenSuccsses() {

	c := SetupControllerCase(http.MethodGet, "/health", nil)

	if suite.Assert().NoError(suite.underTest.Health(c.Ctx)) {
		suite.Assert().Equal(healthJson, strings.TrimSpace(c.Res.Body.String()))
	}

}
