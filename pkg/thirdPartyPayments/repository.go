package thirdPartyPayments

import (
	"fmt"
	"strings"
)

type IMoney interface{}

type IPayment interface{}

type IPaymentService interface{
	//Name() string
	CreatePay(IMoney) (IPayment, error)
}

type IPaymentServiceRepository interface {
	Add(name string, service IPaymentService) error
	All() map[string]IPaymentService
	GeyByName(name string) (IPaymentService, bool)
}

type CaseInsensitiveRepo struct {
	m map[string]IPayment
}

func (r CaseInsensitiveRepo) Add(name string, s IPaymentService) error {
	name = strings.ToLower(name)
	if _, alreadyExist := r.m[name]; alreadyExist {
		return fmt.Errorf("name already used")
	}
	r.m[name] = s
	return nil
}


func (r CaseInsensitiveRepo) All() map[string]IPayment {
	return r.m
}

func (r CaseInsensitiveRepo) GetByName(name string) (IPayment, bool) {
	s, has := r.m[strings.ToLower(name)]
	return s, has
}
