package catalog

type Product struct {
	Title    string
	Price    float64
	Category *Category
}

func NewProduct(title string, price float64, category *Category) Product {
	product := Product{}

	//TODO: Add validation to parameters. If not valid throw error

	product.Title = title
	product.Price = price
	product.Category = category

	return product
}
