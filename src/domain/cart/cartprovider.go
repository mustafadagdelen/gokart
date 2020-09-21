package cart

import (
	catalog "gokart/domain/catalog"
	coupon "gokart/domain/coupon"
)

type CartProvider interface {
	GetCartItems() map[catalog.Product]int
	GetCoupons() map[string]coupon.Coupon
	GetCartTotalProducts() int
}
