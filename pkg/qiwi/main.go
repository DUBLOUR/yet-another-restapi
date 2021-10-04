package qiwi

import (
	"yet-another-restapi/pkg/generalApiReader"
)

type Qiwi struct{
	authToken string
	endpoint string
}

func (q Qiwi) CreateLink(money string) (string, error) {

	req, err := generalApiReader.CreateGetRequest(
		q.endpoint,
		map[string]string{
			"expected": money,
		},
		map[string]string{
			"X-Auth": q.authToken,
		},
	)

	if err != nil {
		return "", nil
	}

	var response string
	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return "", err
	}

	return response, nil
}

