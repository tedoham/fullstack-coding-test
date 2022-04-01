package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/tedoham/fullstack-coding-test/domain"
	"github.com/tedoham/fullstack-coding-test/logger"
	"github.com/tedoham/fullstack-coding-test/service"
	"github.com/tedoham/fullstack-coding-test/util"
)

func Start() {

	// wiring
	dbClient, config := getDbClient()
	router := mux.NewRouter()
	productRepositoryDb := domain.NewProductRepositoryDb(dbClient)
	product := ProductHandlers{service.NewProductService(productRepositoryDb)}

	// define routes
	router.
		HandleFunc("/products", product.GetProducts).
		Methods(http.MethodGet).
		Name("GetProducts")

	router.
		HandleFunc("/product/{product_id:[0-9]+}", product.GetProduct).
		Methods(http.MethodGet).
		Name("GetProduct")

	router.
		HandleFunc("/product/{product_id:[0-9]+}", product.UpdateDeliveryStatus).
		Methods(http.MethodPut).
		Name("UpdateDeliveryStatus")

	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	// starting server
	logger.Info(fmt.Sprintf("Starting server on %s ...", config.SERVER_ADDRESS))
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s", config.SERVER_ADDRESS),
			handlers.CombinedLoggingHandler(os.Stderr, handlers.CORS(headersOK, originsOK, methodsOK)(router))))
}

func getDbClient() (*sqlx.DB, util.Config) {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	client, err := sqlx.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot load connect to database server:", err)
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client, config
}
