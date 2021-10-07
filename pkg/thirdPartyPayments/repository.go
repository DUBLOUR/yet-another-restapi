package thirdPartyPayments

import (
	"fmt"
	"strings"
)

type CaseInsensitiveRepo struct {
	m map[string]*IPaymentService
}

func NewCaseInsensitiveRepo() *CaseInsensitiveRepo {
	return &CaseInsensitiveRepo{
		make(map[string]*IPaymentService),
	}
}

func (r *CaseInsensitiveRepo) Add(name string, s *IPaymentService) error {
	name = strings.ToLower(name)
	if _, exist := r.m[name]; exist {
		return fmt.Errorf("name is already used")
	}
	if r.m == nil {
		r.m = make(map[string]*IPaymentService)
	}
	r.m[name] = s
	return nil
}

func (r CaseInsensitiveRepo) All() map[string]*IPaymentService {
	return r.m
}

func (r CaseInsensitiveRepo) ByName(name string) (*IPaymentService, bool) {
	s, has := r.m[strings.ToLower(name)]
	return s, has
}
