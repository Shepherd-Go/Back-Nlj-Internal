package groups

import (
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/controllers"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/middleware"
	"github.com/labstack/echo/v4"
)

type PaymentMethods interface {
	Resource(g *echo.Group)
}

type paymentmethods struct {
	middlewareJWT      middleware.TokenMiddleware
	paymentMethodsHand controllers.PaymentMethods
}

func NewPaymentMethodsGroup(middlewareJWT middleware.TokenMiddleware, paymentMethodsHand controllers.PaymentMethods) PaymentMethods {
	return &paymentmethods{middlewareJWT, paymentMethodsHand}
}

func (e *paymentmethods) Resource(g *echo.Group) {

	groupPath := g.Group("/employee/methods-payment")

	groupPath.POST("/mobile-payment", e.paymentMethodsHand.RegisterMobilePayment, e.middlewareJWT.Employee)
	groupPath.GET("/mobile-payment/all/:id", e.paymentMethodsHand.SearchAllMobilePayment, e.middlewareJWT.Employee)

}
