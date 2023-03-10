package productmodel

import (
	"fmt"
	"github.com/hvchien216/golang-ecommerce/common"
	"io"
	"strconv"
)

const EntityName = "PRODUCT"

type Product struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Slug            *string       `json:"slug"`
	Code            *string       `json:"code"`
	Image           *string       `json:"image"`
	Images          *string       `json:"images"`
	Price           float64       `json:"price"`
	PriceMin        *float64      `json:"price_min"`
	PriceMax        *float64      `json:"price_max"`
	Status          ProductStatus `json:"status"`
	IsComplete      bool          `json:"is_complete"`
}

type ProductStatus string

const (
	ProductStatusOutOfStock ProductStatus = "OUT_OF_STOCK"
	ProductStatusInStock    ProductStatus = "IN_STOCK"
	ProductStatusRunningLo  ProductStatus = "RUNNING_LO"
)

var AllProductStatus = []ProductStatus{
	ProductStatusOutOfStock,
	ProductStatusInStock,
	ProductStatusRunningLo,
}

func (e ProductStatus) IsValid() bool {
	switch e {
	case ProductStatusOutOfStock, ProductStatusInStock, ProductStatusRunningLo:
		return true
	}
	return false
}

func (e ProductStatus) String() string {
	return string(e)
}

func (e *ProductStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductStatus", str)
	}
	return nil
}

func (e ProductStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
