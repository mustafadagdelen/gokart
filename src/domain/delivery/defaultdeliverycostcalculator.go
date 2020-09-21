package delivery

import (
	"errors"
	"gokart/domain/cart"
)

type DefaultDeliveryCostCalculator struct {
	CostPerDelivery float64
	CostPerProduct  float64
}

const fixedCost float64 = float64(5)

func (d DefaultDeliveryCostCalculator) Calculate(s cart.CartProvider, deliveryCount int) (float64, error) {

	if s == nil {
		return 0, errors.New("Shopping cart cannot be null")
	}

	if deliveryCount <= 0 {
		return 0, errors.New("deliveryCount cannot be less or equals to zero")
	}

	totalProduct := s.GetCartTotalProducts()

	minDeliveryCost := float64(0)
	if totalProduct <= 0 {
		return minDeliveryCost, nil
	}

	val := (d.CostPerDelivery * float64(deliveryCount)) + (d.CostPerProduct * float64(totalProduct)) + fixedCost

	return val, nil
}
