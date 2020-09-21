package cart

import (
	catalog "gokart/domain/catalog"
	coupon "gokart/domain/coupon"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

var shoppingCart ShoppingCart

func setUp() {
	shoppingCart = ShoppingCart{
		cartItems: make(map[catalog.Product]int),
		coupons:   make(map[string]coupon.Coupon),
	}
}

func tearDown() {
	clearShoppingCart()
}

func clearShoppingCart() {
	for key := range shoppingCart.coupons {
		delete(shoppingCart.coupons, key)
	}

	for key := range shoppingCart.cartItems {
		delete(shoppingCart.cartItems, key)
	}
}

func TestAddProductForValidParameters(t *testing.T) {
	var (
		category = catalog.NewCategory("Story Books")
		product  = catalog.NewProduct("Gulliver`s Travels", 34.4, category)
		quantity = 10
	)
	cartItemCount, err := shoppingCart.AddProduct(product, quantity)

	if err != nil {
		t.Error(err)
	}

	if cartItemCount != quantity {
		t.Errorf("Cart item count should be equals to quantity when cart is empty. Quantity : %v, CartItem Count %v", quantity, cartItemCount)
	}
}

func TestAddProductForInvalidParameters(t *testing.T) {
	var (
		category = catalog.NewCategory("")
		product  = catalog.NewProduct("", 0, category)
		quantity = 10
	)
	cartItemCount, err := shoppingCart.AddProduct(product, quantity)

	if err == nil {
		t.Errorf("Shopping cart should have thrown due to empty parameters")
	}

	if cartItemCount == quantity {
		t.Errorf("Cart item count should be equals to quantity when cart is empty. Quantity : %v, CartItem Count %v", quantity, cartItemCount)
	}
}

func TestAddSameProductForQuantity(t *testing.T) {
	var (
		category = catalog.NewCategory("Story Books")
		product  = catalog.NewProduct("Gulliver`s Travels", 34.4, category)
		quantity = 10
	)

	cartItemCount, err := shoppingCart.AddProduct(product, quantity)
	if err != nil {
		t.Error(err)
	}

	cartItemCount, err = shoppingCart.AddProduct(product, quantity)
	if err != nil {
		t.Error(err)
	}

	expected := quantity * 2
	if cartItemCount != expected {
		t.Errorf("Cart item count should be equals to quantity when cart is empty. Quantity : %v, CartItem Count %v", quantity, cartItemCount)
	}
}

func TestAddAmountCoupon(t *testing.T) {
	clearShoppingCart()
	createSampleData()
	var (
		amountCoupon = coupon.AmountCoupon{MinPurchaseAmount: 100, Amount: 20, CouponCode: "TEST-AMOUNT-COUPON", FinishDate: time.Now().AddDate(0, 0, 10)}
	)

	result, err := shoppingCart.AddCoupon(amountCoupon)

	if result == false && err != nil {
		t.Error(err)
	}

	couponCount := len(shoppingCart.coupons)

	expected := 1
	if couponCount < expected {
		t.Error("Coupon count less then expected")
	}
}
func TestAddRateCoupon(t *testing.T) {
	clearShoppingCart()
	createSampleData()
	var (
		rateCoupon = coupon.RateCoupon{MinPurchaseAmount: 100, Rate: 20, CouponCode: "TEST-RATE-COUPON", FinishDate: time.Now().AddDate(0, 0, 10)}
	)

	result, err := shoppingCart.AddCoupon(rateCoupon)

	if result == false && err != nil {
		t.Error(err)
	}

	couponCount := len(shoppingCart.coupons)

	expected := 1
	if couponCount < expected {
		t.Error("Coupon count less then expected")
	}
}

func TestAddMixedCoupon(t *testing.T) {
	clearShoppingCart()
	createSampleData()
	var (
		minPurchaseAmount = float64(100)
		amount            = float64(20)
		rate              = 20
		finishDate        = time.Now().AddDate(0, 0, 10)
		amountCoupon      = coupon.AmountCoupon{MinPurchaseAmount: minPurchaseAmount, Amount: amount, CouponCode: "TEST-AMOUNT-COUPON", FinishDate: finishDate}
		rateCoupon        = coupon.RateCoupon{MinPurchaseAmount: minPurchaseAmount, Rate: rate, CouponCode: "TEST-RATE-COUPON", FinishDate: finishDate}
	)

	rateCouponAddResult, err := shoppingCart.AddCoupon(rateCoupon)

	if rateCouponAddResult == false && err != nil {
		t.Error(err)
	}

	amountCouponAddResult, err := shoppingCart.AddCoupon(amountCoupon)

	if amountCouponAddResult == false && err != nil {
		t.Error(err)
	}

	couponCount := len(shoppingCart.coupons)

	expected := 2
	if couponCount < expected {
		t.Error("Coupon count less then expected")
	}
}
func TestAddSameCoupon(t *testing.T) {
	clearShoppingCart()
	createSampleData()
	var (
		minPurchaseAmount = float64(100)
		amount            = float64(20)
		rate              = 20
		finishDate        = time.Now().AddDate(0, 0, 10)
		amountCoupon      = coupon.AmountCoupon{MinPurchaseAmount: minPurchaseAmount, Amount: amount, CouponCode: "TEST-AMOUNT-COUPON", FinishDate: finishDate}
		rateCoupon        = coupon.RateCoupon{MinPurchaseAmount: minPurchaseAmount, Rate: rate, CouponCode: "TEST-RATE-COUPON", FinishDate: finishDate}
	)

	rateCouponAddResult, err := shoppingCart.AddCoupon(rateCoupon)

	if rateCouponAddResult == false && err != nil {
		t.Error(err)
	}

	amountCouponAddResult, err := shoppingCart.AddCoupon(amountCoupon)

	if amountCouponAddResult == false && err != nil {
		t.Error(err)
	}

	couponCount := len(shoppingCart.coupons)

	expected := 1
	if couponCount < expected {
		t.Error("Coupon count less then expected")
	}
}

func TestCalculateCartTotalPrice(t *testing.T) {
	clearShoppingCart()
	var (
		schoolCategory    = catalog.NewCategory("School")
		bookCategory      = catalog.NewCategory("Books")
		storyBookCategory = catalog.NewCategory("Story Books")
		product1          = catalog.NewProduct("Gulliver`s Travels", 34.4, storyBookCategory)
		quantity1         = 10

		smartDeviceCategory = catalog.NewCategory("Smart Devices")
		laptopCategory      = catalog.NewCategory("Laptop")
		product2            = catalog.NewProduct("Lenovo ThinkPad ", 7500, laptopCategory)
		product3            = catalog.NewProduct("Asus  ", 8000, laptopCategory)
		product4            = catalog.NewProduct("Hp  ", 8000, laptopCategory)
		quantity2           = 2
	)

	bookCategory.SetParentCategory(schoolCategory)
	storyBookCategory.SetParentCategory(bookCategory)
	laptopCategory.SetParentCategory(smartDeviceCategory)

	shoppingCart.AddProduct(product1, quantity1)
	shoppingCart.AddProduct(product2, quantity2)
	shoppingCart.AddProduct(product3, quantity2)
	shoppingCart.AddProduct(product4, quantity2)

	totalPrice := shoppingCart.CalculateCartTotalPriceWithoutDiscounts()

	expected := (product1.Price * float64(quantity1)) + ((product2.Price + product3.Price + product4.Price) * float64(quantity2))

	if totalPrice != expected {
		t.Errorf("Cart price should be %v, But found : %v", expected, totalPrice)
	}

}

func TestGetCartTotalProducts(t *testing.T) {
	clearShoppingCart()
	var (
		category1 = catalog.NewCategory("Story Books")
		product1  = catalog.NewProduct("Gulliver`s Travels", 34.4, category1)
		quantity1 = 10

		category2 = catalog.NewCategory("Animal Books")
		product2  = catalog.NewProduct("Life of Lions", 25, category2)
		quantity2 = 10
	)
	shoppingCart.AddProduct(product1, quantity1)
	shoppingCart.AddProduct(product2, quantity2)

	cartProductCount := shoppingCart.GetCartTotalProducts()
	expected := 2

	if cartProductCount != expected {
		t.Errorf("Cart product count should be %v, But found : %v", expected, cartProductCount)
	}
}

func TestFindAllCategories(t *testing.T) {
	clearShoppingCart()
	createSampleData()
	categories := shoppingCart.FindAllCategories()

	actual := len(categories)
	expected := 5

	if actual != expected {
		t.Errorf("Find All Categories return wrong result. Expected : %v but found %v", expected, actual)
	}
}

func createSampleData() {
	var (
		schoolCategory    = catalog.NewCategory("School")
		bookCategory      = catalog.NewCategory("Books")
		storyBookCategory = catalog.NewCategory("Story Books")
		product1          = catalog.NewProduct("Gulliver`s Travels", 34.4, storyBookCategory)
		quantity1         = 10

		smartDeviceCategory = catalog.NewCategory("Smart Devices")
		laptopCategory      = catalog.NewCategory("Laptop")
		product2            = catalog.NewProduct("Lenovo ThinkPad ", 25, laptopCategory)
		quantity2           = 10
	)

	bookCategory.SetParentCategory(schoolCategory)
	storyBookCategory.SetParentCategory(bookCategory)
	laptopCategory.SetParentCategory(smartDeviceCategory)

	shoppingCart.AddProduct(product1, quantity1)
	shoppingCart.AddProduct(product2, quantity2)
}

func TestFindCategoryProducts(t *testing.T) {
	clearShoppingCart()
	var (
		schoolCategory    = catalog.NewCategory("School")
		bookCategory      = catalog.NewCategory("Books")
		storyBookCategory = catalog.NewCategory("Story Books")
		product1          = catalog.NewProduct("Gulliver`s Travels", 34.4, storyBookCategory)
		quantity1         = 10

		smartDeviceCategory = catalog.NewCategory("Smart Devices")
		laptopCategory      = catalog.NewCategory("Laptop")
		product2            = catalog.NewProduct("Lenovo ThinkPad ", 7500, laptopCategory)
		product3            = catalog.NewProduct("Asus  ", 8000, laptopCategory)
		product4            = catalog.NewProduct("Hp  ", 8000, laptopCategory)
		quantity2           = 10
	)

	bookCategory.SetParentCategory(schoolCategory)
	storyBookCategory.SetParentCategory(bookCategory)
	laptopCategory.SetParentCategory(smartDeviceCategory)

	shoppingCart.AddProduct(product1, quantity1)
	shoppingCart.AddProduct(product2, quantity2)
	shoppingCart.AddProduct(product3, quantity2)
	shoppingCart.AddProduct(product4, quantity2)

	categoryProducts := shoppingCart.FindCategoryProducts(laptopCategory)

	actual := len(categoryProducts)
	expected := 3

	if actual != expected {
		t.Errorf("Find Category Products return wrong result. Expected : %v but found %v", expected, actual)
	}
}

func TestGetCampaignDiscountAmountOfEmptyProducts(t *testing.T) {
	var products []catalog.Product

	var (
		storyBookCategory = catalog.NewCategory("Story Books")
		product1          = catalog.NewProduct("Gulliver`s Travels", 34.4, storyBookCategory)
		quantity          = 1
		laptopCategory    = catalog.NewCategory("Laptop")
		product2          = catalog.NewProduct("Lenovo ThinkPad ", 25, laptopCategory)
	)

	products = append(products, product1)
	products = append(products, product2)

	shoppingCart.AddProduct(product1, quantity)
	shoppingCart.AddProduct(product2, quantity)

	amount := shoppingCart.CalculateAmountOfGivenProducts(products)
	expected := float64(product1.Price) + float64(product2.Price)

	if amount != expected {
		t.Errorf("Method should return zero. But found : %v", amount)
	}
}

func TestGetCampaignDiscountAmount(t *testing.T) {
	var products []catalog.Product
	amount := shoppingCart.CalculateAmountOfGivenProducts(products)
	expected := float64(0)

	if amount != expected {
		t.Errorf("Method should return %v. But found : %v", expected, amount)
	}
}

func TestCalculateAmountOfGivenProducts(t *testing.T) {

}
