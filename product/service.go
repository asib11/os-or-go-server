package product

import "ecommerce/domain"

type service struct {
	prodRepo ProductRepo
}

func NewService(prodRepo ProductRepo) Service {
	return &service{
		prodRepo: prodRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	createdProduct, err := svc.prodRepo.Create(p)
	if err != nil {
		return nil, err
	}

	if createdProduct == nil {
		return nil, nil
	}

	return createdProduct, nil
}

func (svc *service) Get(productID int) (*domain.Product, error) {
	product, err := svc.prodRepo.Get(productID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	return product, nil
}

func (svc *service) List() ([]*domain.Product, error) {
	products, err := svc.prodRepo.List()
	if err != nil {
		return nil, err
	}

	if products == nil {
		return []*domain.Product{}, nil
	}

	return products, nil
}

func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	product, err := svc.prodRepo.Update(p)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	return product, nil
}

func (svc *service) Delete(productID int) error {
	err := svc.prodRepo.Delete(productID)
	if err != nil {
		return err
	}

	return nil
}