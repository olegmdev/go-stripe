package payments

import "errors"
import "fmt"
import stripe "github.com/stripe/stripe-go"
import customer "github.com/stripe/stripe-go/customer"
import product "github.com/stripe/stripe-go/product"
import sku "github.com/stripe/stripe-go/sku"
import order "github.com/stripe/stripe-go/order"
import currency "github.com/stripe/stripe-go/currency"
import orderitem "github.com/stripe/stripe-go/orderitem"

type StripeProvider struct {
  Key string
}

// Create new customer
func (self *StripeProvider) NewCustomer(params *CustomerParams) (*Customer, error) {
  stripe.Key = self.Key

  cardParams := params.Source.Card
  stripeParams := stripe.CustomerParams{
    Balance:       params.Balance,
    Desc:          params.Desc,
    Email:         params.Email,
    BusinessVatID: params.BusinessVatID,
    Source: &stripe.SourceParams{
      Card:     &stripe.CardParams{
        Name:     cardParams.Name,
        Number:   cardParams.Number,
        Month:    cardParams.Month,
        Year:     cardParams.Year,
        Address1: cardParams.Address1,
        Address2: cardParams.Address2,
        City:     cardParams.City,
        State:    cardParams.State,
      },
    },
  }

  stripeCustomer, err := customer.New(&stripeParams)
  return &Customer{
    ID: stripeCustomer.ID,
  }, err
}

// Retrieve a customer by ID
func (self *StripeProvider) GetCustomer(customerId string) (*Customer, error) {
  stripe.Key = self.Key

  stripeCustomer, err := customer.Get(customerId, nil)
  return &Customer{
    ID: stripeCustomer.ID,
  }, err
}

// Retrieve all existing customers
func (self *StripeProvider) GetCustomers() ([]*Customer) {
  stripe.Key = self.Key

  var customers []*Customer

  params := &stripe.CustomerListParams{}
  iterator := customer.List(params)

  for iterator.Next() {
    item := iterator.Customer()
    customers = append(customers, &Customer{
      ID: item.ID,
    })
  }

  return customers
}

// create new product
func (self *StripeProvider) NewProduct(params *ProductParams) (*Product, error) {
  stripe.Key = self.Key

  stripeParams := stripe.ProductParams{
    Name: params.Name,
  }

  stripeProduct, err := product.New(&stripeParams)
  return &Product{
    ID: stripeProduct.ID,
  }, err
}

// create new product with SKU
func (self *StripeProvider) NewProductWithSKU(params *ProductParams, skuParams *SKUParams) (*Product, error) {
  stripe.Key = self.Key

  stripeParams := stripe.ProductParams{
    Name: params.Name,
  }

  stripeSKUParams := stripe.SKUParams{
    Price: skuParams.Price,
    Currency: skuParams.Currency,
    Inventory: stripe.Inventory{
      Type: skuParams.Inventory.Type,
    },
  }

  target, err := product.New(&stripeParams)
  if err != nil {
    return nil, err
  }

  stripeSKUParams.Product = target.ID
  if _, err := sku.New(&stripeSKUParams); err != nil {
    return nil, err
  }

  return &Product{
    ID: target.ID,
  }, nil

}

// retrieve an existing product by ID
func (self *StripeProvider) GetProduct(productId string) (*Product, error) {
  stripe.Key = self.Key

  stripeProduct, err := product.Get(productId)
  return &Product{
    ID: stripeProduct.ID,
  }, err
}

func convertStripeSkus(stripeList *stripe.SKUList) (*SKUList) {
  var skus []*SKU

  for _, sku := range stripeList.Values {
    skus = append(skus, &SKU{
      ID: sku.ID,
    })
  }

  return &SKUList{
    Values: skus,
  }
}

// retrieve all existing products
func (self *StripeProvider) GetProducts() ([]*Product) {
  stripe.Key = self.Key

  var products []*Product

  params := &stripe.ProductListParams{}
  iterator := product.List(params)

  for iterator.Next() {
    item := iterator.Product()
    products = append(products, &Product{
      ID: item.ID,
      Skus: convertStripeSkus(item.Skus),
    })
  }

  return products
}

func convertOrderItemsToStripe(orderItems []*OrderItemParams) ([]*stripe.OrderItemParams) {
  var items []*stripe.OrderItemParams

  for _, item := range orderItems {
    items = append(items, &stripe.OrderItemParams{
      Type: orderitem.SKU,
      Parent: item.Parent,
      Description: item.Description,
    })
  }

  return items
}

// create a new order
func (self *StripeProvider) NewOrder(skuItems []*OrderItemParams, params *OrderParams) (*Order, error) {
  stripe.Key = self.Key

  addressParams := params.Shipping.Address
  stripeParams := stripe.OrderParams{
    Currency: currency.USD,
    Items: convertOrderItemsToStripe(skuItems),
    Shipping: &stripe.ShippingParams{
      Name: params.Shipping.Name,
      Address: &stripe.AddressParams{
        Line1: addressParams.Line1,
        City: addressParams.City,
        Country: addressParams.Country,
        PostalCode: addressParams.PostalCode,
      },
    },
  }

  stripeOrder, err := order.New(&stripeParams)
  return &Order{
    ID: stripeOrder.ID,
  }, err
}

// pay an order
func (self *StripeProvider) PayOrder(orderId string, params *OrderPayParams) (*Order, error) {
  stripe.Key = self.Key

  stripeParams := stripe.OrderPayParams{
    Customer: params.Customer,
  }

  target, err := order.Pay(orderId, &stripeParams)

  if target.Status != stripe.StatusPaid {
    return nil, errors.New(fmt.Sprintf("Order status not set to paid: %v", target.Status))
	}

  return &Order{
    ID: target.ID,
  }, err
}
