package paypal

import (
	"yet-another-restapi/pkg/generalApiReader"
)

type PayPal struct {
	authToken string
	endpoint  string
}

func New(token, endpoint string) *PayPal {
	return &PayPal{
		token,
		endpoint,
	}
}

func (p PayPal) CreateLink(money string) (string, error) {
	req, err := generalApiReader.CreateGetRequest(
		p.endpoint,
		map[string]string{
			"token": p.authToken,
			"money": money,
		},
		map[string]string{},
	)

	if err != nil {
		return "", err
	}

	response := new(struct {
		Id   string `json:"pay_id"`
		Href string `json:"href"`
	})

	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return "", err
	}

	return response.Href, nil
}
