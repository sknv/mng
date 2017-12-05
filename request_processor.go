package mng

import (
	"net/http"

	"github.com/globalsign/mgo"

	"github.com/sknv/mng/middleware"
)

// GetMgoSessionForRequest returns a Mongo session from a request context.
func GetMgoSessionForRequest(r *http.Request) *mgo.Session {
	return r.Context().Value(middleware.CtxKeyMgoSession).(*mgo.Session)
}
