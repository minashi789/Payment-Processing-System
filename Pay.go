package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string, paymentDetails map[string]string) error
}

type PayPalProcessor struct{ logger *log.Logger }

type StripeProcessor struct{ logger *log.Logger }

type BankTransferProcessor struct{ logger *log.Logger }

func Log() (*log.Logger, *os.File) {
	logFile, err := os.OpenFile("payment.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	logger := log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	return logger, logFile
}

func NewPayPalProcessor(logger *log.Logger) *PayPalProcessor {
	return &PayPalProcessor{logger: logger}
}

func NewStripeProcessor(logger *log.Logger) *StripeProcessor {
	return &StripeProcessor{logger: logger}
}

func NewBankTransferProcessor(logger *log.Logger) *BankTransferProcessor {
	return &BankTransferProcessor{logger: logger}
}

func (p *PayPalProcessor) ProcessPayment(amount float64, currency string, paymentDetails map[string]string) error {

	p.logger.Printf("Processing payment via PayPal... Amount: %.2f %s, Email: %s\n", amount, currency, paymentDetails["email"])

	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	email, ok := paymentDetails["email"]
	if !ok || email == "" {
		err := errors.New("Missing email for PayPal payment")
		p.logger.Printf("Error: %s\n", err)
		return err
	}

	if rand.Float64() < 0.8 {
		p.logger.Printf("PayPal payment successful: %.2f %s to %s\n", amount, currency, email)
		return nil
	} else {
		err := errors.New("PayPal payment failed due to insufficient funds or other reasons")
		p.logger.Printf("Error: %s\n", err)
		return err
	}
}

func (s *StripeProcessor) ProcessPayment(amount float64, currency string, paymentDetails map[string]string) error {

	s.logger.Printf("Processing payment via Stripe... Amount: %.2f %s, Card Number: %s\n", amount, currency, paymentDetails["cardNumber"])

	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	cardNumber, ok := paymentDetails["cardNumber"]
	if !ok || cardNumber == "" {
		err := errors.New("Missing card number for Stripe payment")
		s.logger.Printf("Error: %s\n", err)
		return err
	}
	if rand.Float64() < 0.8 {
		s.logger.Printf("Stripe payment successful: %.2f %s to %s\n", amount, currency, cardNumber)
		return nil
	} else {
		err := errors.New("Stripe payment failed due to insufficient funds or other reasons")
		s.logger.Printf("Error: %s\n", err)
		return err
	}

}

func (b *BankTransferProcessor) ProcessPayment(amount float64, currency string, paymentDetails map[string]string) error {

	b.logger.Printf("Processing payment via Bank Transfer... Amount: %.2f %s, Account: %s, Routing: %s\n", amount, currency, paymentDetails["accountNumber"], paymentDetails["routingNumber"])
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	accountNumber, ok := paymentDetails["accountNumber"]
	if !ok || accountNumber == "" {
		err := errors.New("Missing account number for Bank Transfer payment")
		b.logger.Printf("Error: %s\n", err)
		return err
	}

	routingNumber, ok := paymentDetails["routingNumber"]
	if !ok || routingNumber == "" {
		err := errors.New("Missing routing number for Bank Transfer payment")
		b.logger.Printf("Error: %s\n", err)
		return err
	}

	if rand.Float64() < 0.5 {
		b.logger.Printf("Bank Transfer payment successful: %.2f %s to %s\n", amount, currency, accountNumber)
		return nil
	} else {
		err := errors.New("Bank Transfer payment failed due to insufficient funds or other reasons")
		b.logger.Printf("Error: %s\n", err)
		return err
	}
}

func MakePayment(logger *log.Logger, processor PaymentProcessor, amount float64, currency string, paymentDetails map[string]string) {
	if paymentDetails == nil {
		logger.Println("Error: Payment details are missing")
		return
	}
	err := processor.ProcessPayment(amount, currency, paymentDetails)
	if err != nil {
		logger.Printf("Error processing payment: %s\n", err)
	} else {
		logger.Println("Payment processed successfully")
	}
}

func main() {


	logger, logFile := Log()
	defer logFile.Close()

	paypal := NewPayPalProcessor(logger)
	stripe := NewStripeProcessor(logger)
	bank := NewBankTransferProcessor(logger)

	paypalDetails := map[string]string{"email": "test.99@mail.ru"}
	stripeDetails := map[string]string{"cardNumber": "1232131231231643"}
	bankDetails := map[string]string{"accountNumber": "923542523", "routingNumber": "1672734534"}

	MakePayment(logger, paypal, 100.0, "USD", paypalDetails)
	MakePayment(logger, stripe, 200.0, "EUR", stripeDetails)
	MakePayment(logger, bank, 300.0, "RUB", bankDetails)

}
