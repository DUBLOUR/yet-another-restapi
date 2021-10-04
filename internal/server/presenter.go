package server

import "encoding/json"

type JsonPresenter struct{}

func (JsonPresenter) Format(data interface{}) (string, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
