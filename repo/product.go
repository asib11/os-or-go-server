package repo

import "github.com/jmoiron/sqlx"

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(productID int) error
}

type productRepo struct {
	db *sqlx.DB
}

// NewProductRepo creates a new instance of ProductRepo with initial products.
// constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	QUERY := `INSERT INTO products (title, description, price, image_url) VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(QUERY, p.Title, p.Description, p.Price, p.ImageURL)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
	
}

func (r *productRepo) Get(productID int) (*Product, error) {
	QUERY := `SELECT * FROM products WHERE id = $1 LIMIT 1`

	var product Product

	err := r.db.Get(&product, QUERY, productID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) List() ([]*Product, error) {
	QUERY := `SELECT * FROM products`

	var products []*Product

	err := r.db.Select(&products, QUERY)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	QUERY := `UPDATE products SET title = $1, description = $2, price = $3, image_url = $4 WHERE id = $5 RETURNING *`

	var updatedProduct Product

	err := r.db.Get(&updatedProduct, QUERY, p.Title, p.Description, p.Price, p.ImageURL, p.ID)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func (r *productRepo) Delete(productID int) error {
	QUERY := `DELETE FROM products WHERE id = $1`

	_, err := r.db.Exec(QUERY, productID)
	if err != nil {
		return err
	}

	return nil
}
