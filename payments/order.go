package payments

import "fmt"
import "errors"
import stripe "github.com/stripe/stripe-go"
import order "github.com/stripe/stripe-go/order"

func (conf *Config) CreateOrder(params *stripe.OrderParams) (*stripe.Order, error) {
  stripe.Key = conf.Key

  return order.New(params)
}

func (conf *Config) PayOrder(id string, params *stripe.OrderPayParams) (*stripe.Order, error) {
  stripe.Key = conf.Key

  target, err := order.Pay(id, params)

  if target.Status != stripe.StatusPaid {
    return nil, errors.New(fmt.Sprintf("Order status not set to paid: %v", target.Status))
	}

  return target, err
}

// Pay specific order by a card number
func (conf *Config) PayOrderWithCard(id string, params *stripe.OrderPayParams, cardParams *stripe.CardParams) (*stripe.Order, error) {
  stripe.Key = conf.Key

  params.SetSource(cardParams)
  target, err := order.Pay(id, params)

  if target.Status != stripe.StatusPaid {
    return nil, errors.New(fmt.Sprintf("Order status not set to paid: %v", target.Status))
	}

  return target, err
}
