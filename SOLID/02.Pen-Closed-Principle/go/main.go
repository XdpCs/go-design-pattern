package main

import "fmt"

// @Title        main.go
// @Description
// @Create       XdpCs 2024-05-13 15:03
// @Update       XdpCs 2024-05-13 15:03

// Before Open Closed Principle

type Order struct {
	totalCost float32
	payType   string
}

func (o *Order) pay() {
	switch o.payType {
	case "creditCard":
		o.payByCreditCard(o.totalCost)
	case "bitcoin":
		o.payByBitcoin(o.totalCost)
	case "paypal":
		o.payByPaypal(o.totalCost)
	}
}

func (o *Order) payByCreditCard(amount float32) {
	// Pay by credit card
	fmt.Println("Pay by credit card: ", amount)
}

func (o *Order) payByBitcoin(amount float32) {
	// Pay by bitcoin
	fmt.Println("Pay by bitcoin: ", amount)
}

func (o *Order) payByPaypal(amount float32) {
	// Pay by paypal
	fmt.Println("Pay by paypal: ", amount)
}

// After Open Closed Principle

type ModifyOrder struct {
	totalCost float32
	payType   PaymentStrategy
}

type PaymentStrategy interface {
	Pay(amount float32)
}

func (m *ModifyOrder) pay() {
	m.payType.Pay(m.totalCost)
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float32) {
	fmt.Println("Pay by credit card: ", amount)
}

type Bitcoin struct{}

func (b *Bitcoin) Pay(amount float32) {
	fmt.Println("Pay by bitcoin: ", amount)
}

type Paypal struct{}

func (p *Paypal) Pay(amount float32) {
	fmt.Println("Pay by paypal: ", amount)
}

func main() {
	fmt.Println("Before Open Closed Principle")
	o := Order{totalCost: 100, payType: "creditCard"}
	o.pay()
	fmt.Println("After Open Closed Principle")
	m := ModifyOrder{totalCost: 100, payType: &Bitcoin{}}
	m.pay()
}
