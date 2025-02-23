package main

import "fmt"

type paymenter interface {
	pay(amount float32) // you can give return type also
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

// func (p payment) makePayment(amount float32) {
// 	razorPaymentGw := razorPay{}
// 	razorPaymentGw.pay(amount)
// }

type razorPay struct{}

func (r razorPay) pay(amount float32) {
	fmt.Println("Makng Payment Usng RazorPay", amount)
}
func (r razorPay) refund(amount float32, account string) {
	fmt.Println("ReFund Payment Usng RazorPay", amount, account)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Makng Payment Usng stripe", amount)
}

type fakePay struct{}

func (f fakePay) pay(amount float32) {
	fmt.Println("Makng Payment Usng fakePay", amount)
}

func main() {
	razorPaymentGw := razorPay{}
	newPayment := payment{
		gateway: razorPaymentGw,
	}
	newPayment.gateway.pay(200)
	newPayment.gateway.refund(150, "AVISH2034")
}
