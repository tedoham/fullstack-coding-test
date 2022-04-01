package domain

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/tedoham/fullstack-coding-test/errs"
	"github.com/tedoham/fullstack-coding-test/logger"
)

type ProductRepositoryDb struct {
	client *sqlx.DB
}

func NewProductRepositoryDb(dbClient *sqlx.DB) ProductRepositoryDb {
	return ProductRepositoryDb{dbClient}
}

func (p ProductRepositoryDb) GetAllProducts(limit int, offset int) ([]ProductList, *errs.AppError) {
	var err error
	products := make([]ProductList, 0)

	findAllSql := fmt.Sprintf(`
		SELECT 
		products.id as id,
		product_types.name as product_type_name, 
		customers.name as customer_name, 
		delivery_statuses.name as delivery_status, 
		delivery_address, 
		estimated_delivery_date 
		FROM products
		LEFT JOIN product_types on products.product_type_id = product_types.id 
		LEFT JOIN customers on products.customer_id = customers.id
		LEFT JOIN delivery_statuses on products.delivery_status = delivery_statuses.id 
		LIMIT %d
		OFFSET %d
		`, limit, offset)

	err = p.client.Select(&products, findAllSql)

	if err != nil {
		logger.Error("Error while querying product table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return products, nil
}

func (p ProductRepositoryDb) GetProductById(productId int) (*ProductList, *errs.AppError) {
	sqlGetProduct := `
		SELECT 
		products.id as id,
		product_types.name as product_type_name, 
		customers.name as customer_name, 
		delivery_statuses.name as delivery_status, 
		delivery_address, 
		estimated_delivery_date 
		FROM products 
		INNER JOIN product_types on products.product_type_id = product_types.id 
		INNER JOIN customers on products.customer_id = customers.id
		INNER JOIN delivery_statuses on products.delivery_status = delivery_statuses.id 
		WHERE products.id = $1
		LIMIT 1
	`

	var product ProductList
	err := p.client.Get(&product, sqlGetProduct, productId)

	if err != nil {
		logger.Error("Error while fetching product information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &product, nil
}

func (p ProductRepositoryDb) UpdateDeliveryStatus(productId int, deliveryStatusId int) *errs.AppError {

	updateDeliveryStatusSql := `UPDATE products SET delivery_status = $2 WHERE id = $1`

	result, err := p.client.Exec(updateDeliveryStatusSql, productId, deliveryStatusId)
	if err != nil {
		logger.Error("Error while updating product table " + err.Error())
		return errs.NewUnexpectedError("Unexpected error from the database")
	}

	id, err := result.RowsAffected()
	fmt.Println("========vvv=========", id)
	if err != nil {
		logger.Error("Error while getting updated id for product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected error from the database")
	}

	return nil
}
