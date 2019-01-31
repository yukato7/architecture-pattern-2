package external_api

import (
	"context"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"log"
	"os"
)
type Client struct {
	*omise.Client
}

func NewClient() (*Client, error) {
	oc, e := omise.NewClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECREAT_KEY"))
	if e != nil {
		log.Fatal(e)
	}
	client := &Client{oc}
	return client, e
}

func ChargeMoney(ctx context.Context, chargeInfo *operations.CreateCharge) error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	charge, createCharge := &omise.Charge{}, chargeInfo
	if e := client.Do(charge, createCharge); e != nil {
		return err
	}
	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
	return nil
}