package catalog

import "testing"

func TestNewProduct(t *testing.T) {
	categoryTitle := "Book"
	category := NewCategory(categoryTitle)
	title := "Lean Startup"
	price := float64(20)

	product := NewProduct(title, price, category)

	if product.Title != title || product.Price != price || product.Category.title != categoryTitle {
		t.Error("Error on Product.NewProduct")
	}
}
