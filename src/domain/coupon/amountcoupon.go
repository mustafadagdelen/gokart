package coupon

import "time"

type AmountCoupon struct {
	MinPurchaseAmount float64
	Amount            float64
	CouponCode        string
	FinishDate        time.Time
}

func (amountCoupon AmountCoupon) IsApplicable(totalCartAmount float64) bool {
	return totalCartAmount >= amountCoupon.MinPurchaseAmount && amountCoupon.FinishDate.After(time.Now())
}

func (amountCoupon AmountCoupon) ApplyDiscount(totalCartAmount float64) float64 {
	if amountCoupon.IsApplicable(totalCartAmount) {
		return totalCartAmount - amountCoupon.Amount
	}
	return totalCartAmount
}

func (amountCoupon AmountCoupon) GetCouponCode() string {
	return amountCoupon.CouponCode
}
