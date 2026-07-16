package repo


type Product struct {
	ID          int     `db:"id"`
	Title	   string  `db:"title"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	ImageURL    string  `db:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(productID int) error
}

type productRepo struct {
	productList []*Product
}

// NewProductRepo creates a new instance of ProductRepo with initial products.
// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)

	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, p := range r.productList {
		if p.ID == productID {
			return p, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	for i, prod := range r.productList {
		if prod.ID == p.ID {
			r.productList[i] = &p
			return &p, nil
		}
	}
	return nil, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID == productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}

func generateInitialProducts(r *productRepo) {

	products := []*Product{
		{ID: 1, Title: "Sample Product 1", Description: "This is a sample product.", Price: 19.99, ImageURL: "https://example.com/image1.jpg"},
		{ID: 2, Title: "Sample Product 2", Description: "This is another sample product.", Price: 29.99, ImageURL: "https://example.com/image2.jpg"},
		{ID: 3, Title: "Sample Product 3", Description: "This is yet another sample product.", Price: 39.99, ImageURL: "https://example.com/image3.jpg"},
		{ID: 4, Title: "Sample Product 4", Description: "This is a fourth sample product.", Price: 49.99, ImageURL: "https://example.com/image4.jpg"},
		{ID: 5, Title: "Sample Product 5", Description: "This is a fifth sample product.", Price: 59.99, ImageURL: "https://example.com/image5.jpg"},
	}

	for _, p := range products {
		r.productList = append(r.productList, p)
	}
}