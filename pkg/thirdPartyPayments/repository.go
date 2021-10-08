package thirdPartyPayments

import (
	"fmt"
	"strings"
)

type CaseInsensitiveRepo struct {
	m map[string]*IPaymentGateway
}

func NewCaseInsensitiveRepo() *CaseInsensitiveRepo {
	return &CaseInsensitiveRepo{
		make(map[string]*IPaymentGateway),
	}
}

func (r *CaseInsensitiveRepo) Add(name string, s *IPaymentGateway) error {
	name = strings.ToLower(name)
	if _, exist := r.m[name]; exist {
		return fmt.Errorf("name is already used")
	}
	if r.m == nil {
		r.m = make(map[string]*IPaymentGateway)
	}
	r.m[name] = s
	return nil
}

func (r CaseInsensitiveRepo) All() map[string]*IPaymentGateway {
	return r.m
}

func (r CaseInsensitiveRepo) ByName(name string) (*IPaymentGateway, bool) {
	s, has := r.m[strings.ToLower(name)]
	return s, has
}
