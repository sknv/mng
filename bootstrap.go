package mng

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"

	"github.com/sknv/mng/middleware"
)

// BootstrapRouter puts a Mongo session to a request context.
func BootstrapRouter(r chi.Router, session *mgo.Session) {
	r.Use(middleware.WithMgoSession(session))
}
