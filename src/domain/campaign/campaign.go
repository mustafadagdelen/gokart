package campaign

type Campaign interface {
	ApplyDiscount(totalAmount float64, productQuantity int) float64
	GetDiscountAmount(totalAmount float64, productQuantity int) float64
	IsApplicable(productQuantity int) bool
	GetCampaignCode() string
}
