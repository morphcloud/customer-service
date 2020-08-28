package routes

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	"github.com/morphcloud/customer-service/internal/app/handlers/http/rest_v1"
	"github.com/morphcloud/customer-service/internal/app/repositories"
	"github.com/morphcloud/customer-service/internal/app/services"
	"github.com/morphcloud/customer-service/internal/diagnostics"
)

func MapURLPathsToHandlers(r *mux.Router, c *sql.DB, l *log.Logger) {
	r.HandleFunc("/healthz", diagnostics.LivenessHandler(l))
	r.HandleFunc("/readyz", diagnostics.ReadinessHandler(l))

	customerRepository := repositories.NewPSQLCustomerRepository(c)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := rest_v1.NewCustomerHandler(customerService)

	v1Customers := r.PathPrefix("/v1/customers").Subrouter()
	v1Customers.Methods("GET").Path("/{customer_id}").HandlerFunc(customerHandler.FindOne)
}
