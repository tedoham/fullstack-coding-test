package domain

import (
	"reflect"
	"testing"

	"github.com/tedoham/fullstack-coding-test/dto"
)

func TestProductList_ToDto(t *testing.T) {
	tests := []struct {
		name string
		p    ProductList
		want dto.ProductResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ToDto(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductList.ToDto() = %v, want %v", got, tt.want)
			}
		})
	}
}
