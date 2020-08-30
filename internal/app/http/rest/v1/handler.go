package v1

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"golang.org/x/crypto/bcrypt"

	"github.com/morphcloud/customer-service/internal/app/service"
)

type CustomerHandler interface {
	FindOne(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
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

//easyjson:json
type RegistrationRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *customerHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	registrationRequestBody := RegistrationRequestBody{}

	if err = easyjson.Unmarshal(body, &registrationRequestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(registrationRequestBody.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(bytes)
}
