package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/A-Siam/super_delivery/payments_service/data/entities"
	"github.com/A-Siam/super_delivery/payments_service/data/models"
	"github.com/A-Siam/super_delivery/payments_service/ports"
)

func DoPayment(paymentDetails models.PaymentDetails) error {
	fmt.Printf("processing payments by %s on %s with amount %f\n", paymentDetails.Issuer, paymentDetails.Item, paymentDetails.Amount)
	event := entities.Event{
		EventName: paymentEventCreator(paymentDetails),
		Service:   "PAYMENTS_SERVICE",
	}
	db, err := ports.InitEventsDBConnection()
	if err != nil {
		return err
	}
	txn := db.Create(&event)
	if err := txn.Error; err != nil {
		return err
	}
	return nil
}

func RevertPayment(event models.Event) error {
	paymentDetails, err := getPaymentDetails(event.EventName)
	if err != nil {
		return err
	}
	fmt.Printf("reverting payment by %s for item %s with amount %f...\n", paymentDetails.issuer, paymentDetails.item, paymentDetails.amount)
	revertingEvent := entities.Event{
		EventName: paymentEventCreator(models.PaymentDetails{
			Amount: paymentDetails.amount * -1,
			Issuer: paymentDetails.issuer,
			Item:   paymentDetails.item,
		}),
		Service: "PAYMENTS_SERVICE",
	}
	db, err := ports.InitEventsDBConnection()
	if err != nil {
		return err
	}
	txn := db.Create(&revertingEvent)
	if err := txn.Error; err != nil {
		return err
	}
	return nil
}

func getPaymentDetails(s string) (paymentDetails, error) {
	tokens := strings.Split(s, "_")
	issuer, item, amountStr := tokens[2], tokens[4], tokens[6]
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return paymentDetails{}, err
	}
	return paymentDetails{
		issuer: issuer,
		item:   item,
		amount: amount,
	}, nil
}

func paymentEventCreator(details models.PaymentDetails) string {
	return fmt.Sprintf("PAYMENT_BY_%s_ON_%s_AMT_%f", details.Issuer, details.Item, details.Amount)
}

type paymentDetails struct {
	issuer string
	amount float64
	item   string
}
