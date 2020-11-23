package session

import (
	"context"
	"net/http"
	"time"
)

// Based on https://github.com/justinas/nosurf/blob/master/handler.go

type ctxKey int

const (
	sessionKey ctxKey = iota
)

// Config for handling HTTP client cookies
type Config struct {
	// The base cookie that session cookies will be built upon.
	// Set the default session expiration and cookie name here
	BaseCookie http.Cookie
	// Store instance
	Store Store
	// Length of unique session id generated, recommended to be at least 32 bytes
	SessionIDLength int
	// SessionInitFunction for setting IP address or other session life values
	SessionInitFunction func(r *http.Request, session *Session)
}

// DefaultInitFunction called when setting up a new session
func DefaultInitFunction(r *http.Request, session *Session) {
	session.Values["ip"] = r.RemoteAddr
	session.Values["user_agent"] = r.Header.Get("User-Agent")
	session.Values["created"] = time.Now().Unix()
}

// NewConfig with stardard defaults
func NewConfig(store Store) Config {
	return Config{
		BaseCookie: http.Cookie{
			Name:     "session",
			HttpOnly: true,
			// Secure:   true, // TODO would like to enable this, but unit tests and development are harder...
		},
		Store:               store,
		SessionIDLength:     32,
		SessionInitFunction: DefaultInitFunction,
	}
}

func Init(config Config) func(http.ResponseWriter, *http.Request) *Session {
	return func(w http.ResponseWriter, r *http.Request) *Session {

		var session *Session

		// Look for existing, valid session from cookie
		sessionCookie, err := r.Cookie(config.BaseCookie.Name)
		if err == nil {
			if sessionCookie.Value != "" {
				sessionID := DecodeSessionID(sessionCookie.Value)

				// Invalid
				if sessionID != nil && len(sessionID) == config.SessionIDLength {
					var err error
					session, err = config.Store.Get(r.Context(), sessionID)
					if err != nil {
						// TODO
					}
				}
			}
		}

		if session == nil {
			session = NewSession(config.SessionIDLength)

			// Ensure cookie is sent to client this request
			cookie := config.BaseCookie
			cookie.Value = EncodeSessionID(session.ID)
			http.SetCookie(w, &cookie)
		}

		return session
	}
}

// Middleware wrapper for HTTP server
// func Middleware(next http.Handler, config Config) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var session *Session

// 		// Look for existing, valid session from cookie
// 		sessionCookie, err := r.Cookie(config.BaseCookie.Name)
// 		if err == nil {
// 			if sessionCookie.Value != "" {
// 				sessionID := DecodeSessionID(sessionCookie.Value)

// 				// Invalid
// 				if sessionID != nil && len(sessionID) == config.SessionIDLength {
// 					var err error
// 					session, err = config.Store.Get(r.Context(), sessionID)
// 					if err != nil {
// 						// TODO
// 					}
// 				}
// 			}
// 		}

// 		if session == nil {
// 			session = NewSession(config.SessionIDLength)

// 			// Ensure cookie is sent to client this request
// 			cookie := config.BaseCookie
// 			cookie.Value = EncodeSessionID(session.ID)
// 			http.SetCookie(w, &cookie)
// 		}

// 		// Handle request
// 		next.ServeHTTP(w, setSessionContext(r, session))

// 		// Save session changes at end of request
// 		// TODO only if needed
// 		err = config.Store.Save(r.Context(), session)
// 		if err != nil {
// 			// TODO
// 		}

// 	})
// }

// TODO allow multiple session stores
func setSessionContext(r *http.Request, session *Session) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), sessionKey, session))
}

// Load session for user
// func Load(r *http.Request) *Session {
// 	return r.Context().Value(sessionKey).(*Session)
// }
