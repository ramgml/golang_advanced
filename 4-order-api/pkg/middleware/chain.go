package middleware

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

func Chain(middleware ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range middleware {
			next = m(next)
		}
		return next
	}
}