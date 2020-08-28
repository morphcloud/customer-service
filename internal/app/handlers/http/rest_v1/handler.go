package rest_v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"github.com/morphcloud/customer-service/internal/app/services"
)

type CustomerHandler interface {
	FindOne(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	service services.CustomerService
}

func NewCustomerHandler(service services.CustomerService) CustomerHandler {
	return &customerHandler{
		service,
	}
}

func (h *customerHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["customer_id"]
	customer, err := h.service.FindOne(r.Context(), customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(customer, w); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
