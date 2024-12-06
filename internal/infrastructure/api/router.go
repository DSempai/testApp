package api

import (
	"net/http"
	"test/internal/infrastructure/api/user"
	"test/internal/infrastructure/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router struct {
	userHandler *user.Handler
}

func NewRouter(userHandler *user.Handler) *Router {
	return &Router{
		userHandler: userHandler,
	}
}

func (r *Router) Start(addr string) error {
	mux := http.NewServeMux()

	// Register handlers with metrics middleware
	mux.HandleFunc("/users", middleware.MetricsMiddleware(
		"get_user",
		r.userHandler.GetUser,
	))
	mux.HandleFunc("/users/create", middleware.MetricsMiddleware(
		"create_user",
		r.userHandler.CreateUser,
	))

	// Expose metrics endpoint
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, mux)
}
