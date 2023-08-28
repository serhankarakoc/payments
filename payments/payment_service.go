package payments

import "fmt"

type PaymentService struct {
	PayPalProcessor
	StripeProcessor
}

func (ps PaymentService) ProcessPayment(provider string, amount float64, currency string, paymentDetails map[string]interface{}) error {
	var processor PaymentProcessor

	switch provider {
	case "paypal":
		processor = ps.PayPalProcessor
	case "stripe":
		processor = ps.StripeProcessor
	default:
		return fmt.Errorf("unsupported payment provider")
	}

	return processor.ProcessPayment(amount, currency, paymentDetails)
}
