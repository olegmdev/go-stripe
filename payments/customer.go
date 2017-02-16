package payments

import stripe "github.com/stripe/stripe-go"
import customer "github.com/stripe/stripe-go/customer"

// Create new customer
func (conf *Config) CreateCustomer(params *stripe.CustomerParams) (*stripe.Customer, error) {
  stripe.Key = conf.Key

  return customer.New(params)
}

// Retrieve a customer by ID
func (conf *Config) GetCustomer(customerId string) (*stripe.Customer, error) {
  stripe.Key = conf.Key

  return customer.Get(customerId, nil)
}

// Retrieve all existing customers
func (conf *Config) GetCustomers() ([]*stripe.Customer) {
  stripe.Key = conf.Key

  var customers []*stripe.Customer

  params := &stripe.CustomerListParams{}
  iterator := customer.List(params)

  for iterator.Next() {
    item := iterator.Customer()
    customers = append(customers, item)
  }

  return customers
}
