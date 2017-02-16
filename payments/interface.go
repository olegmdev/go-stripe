package payments

import stripe "github.com/stripe/stripe-go"

type Payments interface {
  CreateCustomer(*stripe.CustomerParams) (*stripe.Customer, error)
  GetCustomer(string) (*stripe.Customer, error)
  GetCustomers() ([]*stripe.Customer)

  CreateProduct(*stripe.ProductParams) (*stripe.Product, error)
  CreateProductWithSKU(*stripe.ProductParams, *stripe.SKUParams) (*stripe.Product, error)
  GetProduct(string) (*stripe.Product, error)
  GetProducts() ([]*stripe.Product)

  CreateOrder(*stripe.OrderParams) (*stripe.Order, error)
  PayOrder(string, *stripe.OrderPayParams) (*stripe.Order, error)
  PayOrderWithCard(string, *stripe.OrderPayParams, *stripe.CardParams) (*stripe.Order, error)
}
