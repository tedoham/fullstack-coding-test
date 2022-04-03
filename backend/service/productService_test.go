package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/tedoham/fullstack-coding-test/domain"
	mock "github.com/tedoham/fullstack-coding-test/service/mocks"
	"github.com/tedoham/fullstack-coding-test/util"
)

func Test_GetProducts(t *testing.T) {
	product := randomProduct()

	controller := gomock.NewController(t)
	defer controller.Finish()

	service := mock.NewMockProductService(controller)
	service.EXPECT().GetProduct(1).Return([]domain., nil)

	server := mux.NewRouter()
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/product/%d", product.ProductId)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.ServeHTTP(recorder, request)

	// check reponse
	require.Equal(t, http.StatusOK, recorder.Code)

}

func randomProduct() domain.ProductList {
	return domain.ProductList{
		ProductId:             util.RandomString(10),
		ProductTypeName:       util.RandomString(10),
		CustomerName:          util.RandomString(10),
		DeliveryStatus:        util.RandomString(4),
		DeliveryAddress:       util.RandomString(10),
		EstimatedDeliveryDate: time.Now(),
	}
}
