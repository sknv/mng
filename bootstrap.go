package mng

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"

	"github.com/sknv/mng/middleware"
)

// BootstrapRouter puts a Mongo session to a request context.
func BootstrapRouter(router chi.Router, session *mgo.Session) {
	router.Use(middleware.WithMgoSession(session))
}
