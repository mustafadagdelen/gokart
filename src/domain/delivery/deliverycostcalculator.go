package delivery

import (
	"gokart/domain/cart"
)

type DeliverCostCalculator interface {
	Calculate(s cart.CartProvider) float64
}
