package product_test

import (
	"context"
	"testing"
	"time"

	"castanha/database"
	"castanha/models/product"
)

func TestSelectByType(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := database.GetDB()

	products := make([]product.Product, 0, 10)

	err := db.WithContext(ctx).Where("type = ?", "aaa").Find(&products).Error
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range products {
		t.Log(v.Name)
		t.Log(v.InCashValue)
		t.Log(v.InstallmentValue)
	}
}
