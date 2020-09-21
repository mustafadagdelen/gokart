package coupon

type Coupon interface {
	// CreateCoupon(minPurchaseAmount float64, amount float64) Coupon
	ApplyDiscount(totalAmount float64) float64
	IsApplicable(cost float64) bool
	GetCouponCode() string
}
