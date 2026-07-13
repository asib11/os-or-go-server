package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func GetProductByID(id int) *Product {
	for _, p := range productList {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func UpdateProduct( updatedProduct Product){
	for idx, p := range productList {
		if p.ID == updatedProduct.ID {
			productList[idx] = updatedProduct
		}
	}
}

func DeleteProduct(id int)  {
	var tempList []Product

	for _, p := range productList {
		if p.ID != id {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	productList = []Product{
		{ID: 1, Title: "Product 1", Description: "Description of Product 1", Price: 19.99, ImageUrl: "https://example.com/product1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description of Product 2", Price: 29.99, ImageUrl: "https://example.com/product2.jpg"},
		{ID: 3, Title: "Product 3", Description: "Description of Product 3", Price: 39.99, ImageUrl: "https://example.com/product3.jpg"},
	}
}
