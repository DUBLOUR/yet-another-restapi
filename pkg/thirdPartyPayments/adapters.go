package thirdPartyPayments

import (
	"fmt"
	"math/rand"
	"time"
	"yet-another-restapi/pkg/paypal"
	"yet-another-restapi/pkg/qiwi"
)

const PayPalEndpoint string = "127.0.0.1:8901"
const QiwiEndpoint string = "127.0.0.1:8902"

type PayPal struct {
	adaptee *paypal.PayPal
}

func DefaultPayPal(token string) *PayPal {
	return &PayPal{
		adaptee: paypal.New(token, PayPalEndpoint),
	}
}

func (p PayPal) CreatePay(m IMoney) (IPayment, error) {
	return p.adaptee.CreateLink(fmt.Sprintf("%v", m))
}

type Qiwi struct {
	adaptee *qiwi.Qiwi
}

func DefaultQiwi(token string) *Qiwi {
	return &Qiwi{
		adaptee: qiwi.New(token, QiwiEndpoint),
	}
}

func (p Qiwi) CreatePay(m IMoney) (IPayment, error) {
	return p.adaptee.CreateLink(fmt.Sprintf("%v", m))
}


func randomString(length int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

	rand.Seed(time.Now().UnixNano())
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

type VirtualPay struct {
	formatStr string
}

func DefaultVirtualPay() *VirtualPay {
	return &VirtualPay{
		"https://virtualpay.dev/?money=%v&id=%v",
	}
}

func (p VirtualPay) CreatePay(m IMoney) (IPayment, error) {
	return fmt.Sprintf(p.formatStr, m, randomString(24)), nil
}

func (p PayPal) Name() string     { return "PayPal" }
func (p Qiwi) Name() string       { return "Qiwi" }
func (p VirtualPay) Name() string { return "VirtualPay" }
