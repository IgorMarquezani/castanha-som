package description_test

import (
	"castanha/database"
	"castanha/models/product/description"
	"context"
	"testing"
)

func TestDelete(t *testing.T) {
	db := database.GetDB()

	ctx := context.Background()

	err := db.WithContext(ctx).Where("product_name = ?", "Violão").
		Delete(&description.Description{}).Error

	if err != nil {
		t.Log(err.Error())
	}
}
