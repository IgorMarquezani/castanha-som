package product

import (
	"castanha/models/product/description"
)

type Product struct {
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	InCashValue      float32 `json:"in_cash_value"`
	InstallmentValue float32 `json:"installment_value"`
	ImageName        string

	Descriptions []description.Description `json:"descriptions" gorm:"-"`
}
