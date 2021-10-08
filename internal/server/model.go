package server

import (
	"fmt"
	"yet-another-restapi/internal/product"
	"yet-another-restapi/pkg/thirdPartyPayments"
)

type IPaymentGatewayRepository interface {
	Add(name string, s *thirdPartyPayments.IPaymentGateway) error
	All() map[string]*thirdPartyPayments.IPaymentGateway
	ByName(name string) (*thirdPartyPayments.IPaymentGateway, bool)
}

type IProductsRepository interface {
	LoadJson(filename string) error
	Add(name string, product product.Product) error
	ByName(name string) (product.Product, bool)
}

type Model struct {
	Gateways          IPaymentGatewayRepository
	Products          IProductsRepository
	errorRedirectLink string
}

func DefaultModel() *Model {
	return &Model{
		thirdPartyPayments.NewCaseInsensitiveRepo(),
		product.NewRepo(),
		DefaultErrorRedirectLink,
	}
}

func (m *Model) DefaultResponse() string {
	return m.errorRedirectLink
}

func (m *Model) LoadProducts(jsonFilename string) error {
	return m.Products.LoadJson(jsonFilename)
}

func (m *Model) AddGateway(name string, s *thirdPartyPayments.IPaymentGateway) error {
	return m.Gateways.Add(name, s)
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

func (m *Model) CreateBill(product string, gatewayName string) (IPayment, error) {
	g, has := m.Gateways.ByName(gatewayName)
	if !has {
		return "", fmt.Errorf("unknown payment service")
	}
	cost, err := m.price(product)
	if err != nil {
		return "", err
	}
	bill, err := (*g).CreatePay(cost)
	if err != nil {
		return "", err
	}

	return bill, nil
}
