package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Repository struct {
	m map[string]Product
}

func NewRepo() *Repository {
	return &Repository{
		map[string]Product{},
	}
}

func (r *Repository) LoadJson(filename string) error {
	r.m = make(map[string]Product)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("products file doesn't exist")
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("products file cannot be read")
	}
	err = json.Unmarshal(data, &r.m)
	if err != nil {
		return fmt.Errorf("products file is corrupted")
	}
	return nil
}

func (r Repository) Add(name string, product Product) error {
	r.m[name] = product
	return nil
}

func (r Repository) ByName(name string) (Product, bool) {
	p, has := r.m[name]
	return p, has
}
