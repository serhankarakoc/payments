# payments
Golang Payments Package
# Kullanım
```
package main

import "payments"

func main() {
	service := payments.PaymentService{}

	paypalPaymentDetails := map[string]interface{}{"paypal_param": "value"}
	stripePaymentDetails := map[string]interface{}{"stripe_param": "value"}

	err := service.ProcessPayment("paypal", 100.0, "USD", paypalPaymentDetails)
	if err != nil {
		fmt.Println("PayPal payment failed:", err)
	}

	err = service.ProcessPayment("stripe", 150.0, "EUR", stripePaymentDetails)
	if err != nil {
		fmt.Println("Stripe payment failed:", err)
	}
}
```
