package payments

type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string, paymentDetails map[string]interface{}) error
}
