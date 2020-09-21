package coupon

import (
	"time"
)

type RateCoupon struct {
	MinPurchaseAmount float64
	Rate              int
	CouponCode        string
	FinishDate        time.Time
}

func (rateCoupon RateCoupon) IsApplicable(totalCartAmount float64) bool {
	return totalCartAmount > rateCoupon.MinPurchaseAmount && rateCoupon.FinishDate.After(time.Now())
}

func (rateCoupon RateCoupon) ApplyDiscount(totalCartAmount float64) float64 {
	const maxPercentage float64 = 100
	if rateCoupon.IsApplicable(totalCartAmount) {
		return totalCartAmount - float64(totalCartAmount/100*float64(rateCoupon.Rate))
	}
	return totalCartAmount
}

func (rateCoupon RateCoupon) GetCouponCode() string {
	return rateCoupon.CouponCode
}
