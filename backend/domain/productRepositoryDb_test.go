package domain

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tedoham/fullstack-coding-test/errs"
)

func TestNewProductRepositoryDb(t *testing.T) {
	type args struct {
		dbClient *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want ProductRepositoryDb
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductRepositoryDb(tt.args.dbClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepositoryDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepositoryDb_GetAllProducts(t *testing.T) {
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name  string
		p     ProductRepositoryDb
		args  args
		want  []ProductList
		want1 *errs.AppError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.GetAllProducts(tt.args.limit, tt.args.offset)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDb.GetAllProducts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDb.GetAllProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductRepositoryDb_GetProductById(t *testing.T) {
	type args struct {
		productId int
	}
	tests := []struct {
		name  string
		p     ProductRepositoryDb
		args  args
		want  *ProductList
		want1 *errs.AppError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.GetProductById(tt.args.productId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDb.GetProductById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDb.GetProductById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductRepositoryDb_UpdateDeliveryStatus(t *testing.T) {
	type args struct {
		productId        int
		deliveryStatusId int
	}
	tests := []struct {
		name string
		p    ProductRepositoryDb
		args args
		want *errs.AppError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.UpdateDeliveryStatus(tt.args.productId, tt.args.deliveryStatusId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDb.UpdateDeliveryStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
