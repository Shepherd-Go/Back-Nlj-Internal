package groups

import (
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/controllers"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Resorce(g *echo.Group)
}

type session struct {
	sessionHand controllers.Session
}

func NewGroupSession(sessionHand controllers.Session) Session {
	return &session{sessionHand}
}

func (s *session) Resorce(g *echo.Group) {

	groupPath := g.Group("/session")

	groupPath.POST("", s.sessionHand.Session)

}
