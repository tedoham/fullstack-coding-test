package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tedoham/fullstack-coding-test/dto"
	"github.com/tedoham/fullstack-coding-test/service"
)

type ProductHandlers struct {
	service service.ProductService
}

func (ch ProductHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {

	products, err := ch.service.GetProducts(10, 1)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, products)
	}
}

func (ch ProductHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["product_id"]
	productId, errId := strconv.Atoi(id)
	if errId != nil {
		writeResponse(w, http.StatusBadRequest, errId)
	}

	product, err := ch.service.GetProduct(productId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, product)
	}
}

func (ph ProductHandlers) UpdateDeliveryStatus(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["product_id"]
	convertedId, errId := strconv.Atoi(id)
	if errId != nil {
		writeResponse(w, http.StatusBadRequest, errId)
	}

	var request dto.DeliveryStatusRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		statusId, _ := strconv.Atoi(request.StatusID)
		err := ph.service.UpdateDeliveryStatus(convertedId, statusId)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, "Delivery status updated successfully")
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
