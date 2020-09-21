package cart

import (
	catalog "gokart/domain/catalog"
	coupon "gokart/domain/coupon"
)

type FakeCart struct {
	CartItems map[catalog.Product]int
	Coupons   map[string]coupon.Coupon
}

func (fk FakeCart) GetCartItems() map[catalog.Product]int {
	return fk.CartItems
}

func (fk FakeCart) GetCoupons() map[string]coupon.Coupon {
	return fk.Coupons
}

func (fk FakeCart) GetCartTotalProducts() int {
	return len(fk.CartItems)
}
