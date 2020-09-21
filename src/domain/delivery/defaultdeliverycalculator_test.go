package delivery

import (
	"gokart/domain/cart"
	"gokart/domain/catalog"
	"gokart/domain/coupon"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}

var shoppingCart cart.FakeCart

func setUp() {
	shoppingCart = cart.FakeCart{
		CartItems: make(map[catalog.Product]int),
		Coupons:   make(map[string]coupon.Coupon),
	}
}

func TestCalculateForNullShoppingCart(t *testing.T) {

	costPerDelivery := float64(1.2)
	costPerProduct := float64(0.5)
	d := DefaultDeliveryCostCalculator{costPerDelivery, costPerProduct}
	deliveryCount := 5

	deliveryCost, err := d.Calculate(nil, deliveryCount)

	if err == nil || deliveryCost > 0 {
		t.Error("Delivery cost calculater could thrown error")
	}
}

func TestCalculateForZerDeliveryCount(t *testing.T) {

	costPerDelivery := float64(1.2)
	costPerProduct := float64(0.5)
	deliveryCount := 5

	d := DefaultDeliveryCostCalculator{costPerDelivery, costPerProduct}

	deliveryCost, err := d.Calculate(shoppingCart, deliveryCount)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Delivery cost %v", deliveryCost)
}

func TestCalculateForNoProduct(t *testing.T) {
	costPerDelivery := float64(1.2)
	costPerProduct := float64(0.5)
	d := DefaultDeliveryCostCalculator{costPerDelivery, costPerProduct}
	deliveryCount := 1

	deliveryCost, err := d.Calculate(shoppingCart, deliveryCount)

	if err != nil {
		t.Error(err)
	}

	if deliveryCost != 0 {
		t.Errorf("Delivery cost should return zero when shopping cart has no item. But found : %v", deliveryCost)
	}
}

func TestCalculate(t *testing.T) {
	var (
		category1 = catalog.NewCategory("Story Books")
		product1  = catalog.NewProduct("Gulliver`s Travels", 34.4, category1)
		quantity1 = 10

		category2 = catalog.NewCategory("Laptop")
		product2  = catalog.NewProduct("Lenovo Thinkpad", 10304.4, category2)
		quantity2 = 1
	)

	shoppingCart.CartItems[product1] = quantity1
	shoppingCart.CartItems[product2] = quantity2

	costPerDelivery := float64(1.2)
	costPerProduct := float64(0.5)
	d := DefaultDeliveryCostCalculator{costPerDelivery, costPerProduct}
	deliveryCount := 2

	deliveryCost, err := d.Calculate(shoppingCart, deliveryCount)

	if err != nil {
		t.Error(err)
	}

	fixedCost := float64(5)

	expected := (costPerDelivery * float64(deliveryCount)) + (costPerProduct * float64(shoppingCart.GetCartTotalProducts())) + fixedCost

	if deliveryCost != expected {
		t.Errorf("Delivery cost calculation is wrong. Expected : %v. Delivery Cost : %v ", expected, deliveryCost)
	}

}
