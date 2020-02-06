package middleware

import (
	http2 "github.com/RobyFerro/go-web/app/http"
	"github.com/RobyFerro/go-web/exception"
	"github.com/gorilla/sessions"
	"net/http"
)

// Your custom middleware
func (Middleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := http2.ServiceContainer.Invoke(func(s *sessions.CookieStore) {

			session, err := s.Get(r, "basic-auth")
			if err != nil {
				exception.ProcessError(err)
			}

			if session.Values["user"] == nil {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte("Forbidden!"))
			}

		}); err != nil {
			exception.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}