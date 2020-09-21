package cart

import (
	"errors"
	"gokart/domain/catalog"
	"gokart/domain/coupon"
)

type ShoppingCart struct {
	cartItems map[catalog.Product]int
	coupons   map[string]coupon.Coupon
}

func (sc ShoppingCart) GetCartItems() map[catalog.Product]int {
	return sc.cartItems
}

func (sc ShoppingCart) GetCoupons() map[string]coupon.Coupon {
	return sc.coupons
}

func (sc *ShoppingCart) AddProduct(product catalog.Product, quantity int) (int, error) {
	if quantity <= 0 {
		return 0, errors.New("Quantity ")
	}

	if product == (catalog.Product{}) || product.Title == "" || product.Price <= 0 {
		return 0, errors.New("Product should have title, price and category")
	}

	if _, exist := sc.cartItems[product]; exist {
		sc.cartItems[product] += quantity
	} else {
		sc.cartItems[product] = quantity
	}

	return sc.cartItems[product], nil
}

func (sc *ShoppingCart) AddCoupon(coupon coupon.Coupon) (bool, error) {
	if !coupon.IsApplicable(sc.GetCartAmount()) {
		return false, errors.New("Coupon is not applicable")
	}

	if _, exist := sc.coupons[coupon.GetCouponCode()]; exist {
		return false, errors.New("Coupon already applied")
	}

	sc.coupons[coupon.GetCouponCode()] = coupon

	return true, nil
}

func (sc *ShoppingCart) FindAllCategories() []*catalog.Category {

	if len(sc.cartItems) <= 0 {
		return nil
	}

	categories := []*catalog.Category{}

	for key, _ := range sc.cartItems {
		categories = append(categories, key.Category.GetCategoryTree()...)
	}

	return categories
}

func (sc *ShoppingCart) FindCategoryProducts(category *catalog.Category) []catalog.Product {
	if len(sc.cartItems) <= 0 {
		return nil
	}

	categories := sc.FindAllCategories()

	if len(categories) <= 0 {
		return nil
	}

	products := []catalog.Product{}

	for key, _ := range sc.cartItems {
		if key.Category == category {
			products = append(products, key)
		}
	}

	return products
}

func (sc *ShoppingCart) GetCampaignDiscountAmount() float64 {
	if len(sc.cartItems) <= 0 {
		return float64(0)
	}

	categories := sc.FindAllCategories()

	var totalDiscountPrice float64

	for _, category := range categories {
		products := sc.FindCategoryProducts(category)
		productQuantity := sc.FindTotalProductQuantity(products)

		campaigns := category.GetAvailableCampaigns(productQuantity)

		if len(campaigns) <= 0 {
			continue
		}

		for _, campaign := range campaigns {
			productTotalPrice := sc.CalculateAmountOfGivenProducts(products)
			totalDiscountPrice += campaign.GetDiscountAmount(productTotalPrice, productQuantity)
		}
	}

	return totalDiscountPrice
}

func (sc *ShoppingCart) CalculateAmountOfGivenProducts(products []catalog.Product) float64 {
	if len(products) <= 0 {
		return float64(0)
	}

	var totalPrice float64
	for _, product := range products {
		quantity := sc.cartItems[product]
		totalPrice += product.Price * float64(quantity)
	}

	return totalPrice
}

func (sc *ShoppingCart) GetCouponAppliedPrice() float64 {
	if len(sc.cartItems) <= 0 {
		return float64(0)
	}

	if len(sc.coupons) <= 0 {
		return float64(0)
	}

	price := sc.GetCampaignDiscountAmount()

	for _, coupon := range sc.coupons {
		price = coupon.ApplyDiscount(price)
	}

	return price
}

func (sc *ShoppingCart) FindTotalProductQuantity(products []catalog.Product) int {
	if len(products) <= 0 {
		return 0
	}

	var total int
	for _, product := range products {
		total += sc.cartItems[product]
	}

	return 0
}

func (sc *ShoppingCart) CalculateCartTotalPriceWithoutDiscounts() float64 {
	var totalPrice float64

	for key, value := range sc.cartItems {
		totalPrice += key.Price * float64(value)
	}

	return totalPrice
}

func (sc *ShoppingCart) GetCartAmount() float64 {
	return sc.CalculateCartTotalPriceWithoutDiscounts() - sc.GetCampaignDiscountAmount() - sc.GetCouponAppliedPrice()
}

func (sc ShoppingCart) GetCartTotalProducts() int {
	return len(sc.cartItems)
}
