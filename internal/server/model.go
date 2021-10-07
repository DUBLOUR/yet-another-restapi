package server

import (
	"fmt"
	"yet-another-restapi/internal/product"
	"yet-another-restapi/pkg/thirdPartyPayments"
)

//type IPaymentService interface {
//	CreatePay(thirdPartyPayments.IMoney) (thirdPartyPayments.IPayment, error)
//}

type IPaymentServiceRepository interface {
	Add(name string, s *thirdPartyPayments.IPaymentService) error
	All() map[string]*thirdPartyPayments.IPaymentService
	ByName(name string) (*thirdPartyPayments.IPaymentService, bool)
}

type IProductsRepository interface {
	LoadJson(filename string) error
	Add(name string, product product.Product) error
	ByName(name string) (product.Product, bool)
}

type Model struct {
	Services IPaymentServiceRepository
	Products IProductsRepository
}

func DefaultModel() *Model {
	return &Model{
		thirdPartyPayments.NewCaseInsensitiveRepo(),
		product.NewRepo(),
	}
}

func (m *Model) LoadProducts(jsonFilename string) error {
	return m.Products.LoadJson(jsonFilename)
}

func (m *Model) AddService(name string, s *thirdPartyPayments.IPaymentService) error {
	return m.Services.Add(name, s)
}

func (m Model) price(product string) (product.IMoney, error) {
	p, exist := m.Products.ByName(product)
	if !exist {
		return nil, fmt.Errorf("unknown product")
	}
	if !p.Available {
		return nil, fmt.Errorf("product is not available for sale")
	}
	return p.Price, nil
}

func (m *Model) CreateBill(product string, serviceName string) (IPayment, error) {
	service, has := m.Services.ByName(serviceName)
	if !has {
		return "", fmt.Errorf("unknown payment service")
	}
	cost, err := m.price(product)
	if err != nil {
		return "", err
	}
	bill, err := (*service).CreatePay(cost)
	if err != nil {
		return "", err
	}

	return bill, nil
}
