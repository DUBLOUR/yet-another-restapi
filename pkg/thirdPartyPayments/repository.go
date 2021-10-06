package thirdPartyPayments

import (
	"fmt"
	"strings"
)

type IMoney interface{}

type IPayment interface{}

type IPaymentService interface {
	CreatePay(IMoney) (IPayment, error)
}

type CaseInsensitiveRepo struct {
	m map[string]IPaymentService
}

func (r CaseInsensitiveRepo) Add(name string, s IPaymentService) error {
	name = strings.ToLower(name)
	if _, alreadyExist := r.m[name]; alreadyExist {
		return fmt.Errorf("name already used")
	}
	r.m[name] = s
	return nil
}

func (r CaseInsensitiveRepo) All() map[string]IPaymentService {
	return r.m
}

func (r CaseInsensitiveRepo) ByName(name string) (IPaymentService, bool) {
	s, has := r.m[strings.ToLower(name)]
	return s, has
}
