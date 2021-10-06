package server

import (
	"fmt"
	"yet-another-restapi/pkg/thirdPartyPayments"
)

type IMoney interface{}

type IPaymentService interface {
	Name() string
	CreatePay(thirdPartyPayments.IMoney) (thirdPartyPayments.IPayment, error)
}

type IPaymentServiceRepository interface {
	Add(name string, s thirdPartyPayments.IPaymentService) error
	All() map[string]thirdPartyPayments.IPaymentService
	ByName(name string) (thirdPartyPayments.IPaymentService, bool)
}

type IProduct interface {
	Cost() IMoney
	IsAvailable() bool
}

type IProductsRepository interface {
	Add(name string, product IProduct) error
	ByName(name string) (IProduct, bool)
}

type Model struct {
	services IPaymentServiceRepository
	products IProductsRepository
}

type ProductRepo struct {
	m map[string]IProduct
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{
		make(map[string]IProduct),
	}
}

func (r ProductRepo) Add(name string, product IProduct) error {
	r.m[name] = product
	return nil
}

func (r ProductRepo) ByName(name string) (IProduct, bool) {
	p, has := r.m[name]
	return p, has
}

func DefaultModel() *Model {
	return &Model{
		thirdPartyPayments.CaseInsensitiveRepo{},
		NewProductRepo(),
	}
}

func (m Model) price(product string) (IMoney, error) {
	p, exist := m.products.ByName(product)
	if !exist {
		return nil, fmt.Errorf("unknown product")
	}
	if !p.IsAvailable() {
		return nil, fmt.Errorf("product is not available for sale")
	}
	return p.Cost(), nil
}

func (m Model) CreateBill(product string, serviceName string) (IPayment, error) {
	service, has := m.services.ByName(serviceName)
	if !has {
		return "", fmt.Errorf("unknown payment service")
	}
	cost, err := m.price(product)
	if err != nil {
		return "", err
	}
	bill, err := service.CreatePay(cost)
	if err != nil {
		return "", err
	}

	return bill, nil
}
