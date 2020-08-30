package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	restV1 "github.com/morphcloud/customer-service/internal/app/http/rest/v1"
	"github.com/morphcloud/customer-service/internal/app/repository"
	"github.com/morphcloud/customer-service/internal/app/service"
	"github.com/morphcloud/customer-service/internal/diagnostics"
)

func MapURLPathsToHandlers(r *mux.Router, c *pgx.Conn, l *log.Logger) {
	r.HandleFunc("/healthz", diagnostics.LivenessHandler(l))
	r.HandleFunc("/readyz", diagnostics.ReadinessHandler(l))

	customerRepository := repository.NewPGXCustomerRepository(c)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := restV1.NewCustomerHandler(customerService)

	v1Customers := r.PathPrefix("/v1/customers").Subrouter()
	v1Customers.Methods("GET").Path("/{customer_id}").HandlerFunc(customerHandler.FindOne)
	v1Customers.Methods("POST").HandlerFunc(customerHandler.Register)
}
