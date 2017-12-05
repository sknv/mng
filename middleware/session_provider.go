package middleware

import (
	"context"
	"net/http"

	"github.com/globalsign/mgo"
)

type ctxKey string

const CtxKeyMgoSession = ctxKey("_mgo.Session")

// WithMgoSession puts a Mongo session instance to a request context.
func WithMgoSession(
	session *mgo.Session,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Copy Mongo session and schedule a clean up.
			sessionCopy := session.Copy()
			defer sessionCopy.Close()

			// Put the database into a request context.
			ctx := context.WithValue(r.Context(), CtxKeyMgoSession, sessionCopy)

			// Process request.
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
