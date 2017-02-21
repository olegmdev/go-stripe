package main

import "fmt"
import "wrapper/payments"

type PaymentMaker struct {
  paymentProvider payments.IPaymentProvider
}

func (self *PaymentMaker) doAPayment() {
  // create fabrics
  customerFactory := payments.CustomerFactory{ IPaymentProvider: self.paymentProvider }
  productFactory := payments.ProductFactory{ IPaymentProvider: self.paymentProvider }
  orderFactory := payments.OrderFactory{ IPaymentProvider: self.paymentProvider }

  // create a customer
  customer, _ := customerFactory.New(&payments.CustomerParams{
		Balance:       10000,
		Desc:          "Test Customer",
		Email:         "a@b.com",
		BusinessVatID: "123456789",
    Source:        &payments.SourceParams{
      Card:     &payments.CardParams{
        Name:     "Tester",
        Number:   "4242424242424242",
        Month:    "06",
        Year:     "20",
        Address1: "1234 Main Street",
        Address2: "Apt 1",
        City:     "Anytown",
        State:    "CA",
      },
    },
	})

  // create a product
  productFactory.NewWithSKU(&payments.ProductParams{
    Name: "T-shirt",
  }, &payments.SKUParams{
    Price: 150,
    Currency: "usd",
    Inventory: payments.Inventory{Type: "infinite"},
  })

  // Get all existing products and parse SKUs
  products := productFactory.List()
  var skuItems []*payments.OrderItemParams
  for _, product := range products {
    skuItems = append(skuItems, &payments.OrderItemParams{
      Type:   "sku",
      Parent: product.Skus.Values[0].ID,
      Description: "Some description",
    })
  }

  // create an order
  order, _ := orderFactory.New(skuItems, &payments.OrderParams{
    Shipping: &payments.ShippingParams{
      Name: "Noah Davis",
      Address: &payments.AddressParams{
        Line1: "1234 Main Street",
        City: "San Francisco",
        Country: "US",
        PostalCode: "94111",
      },
    },
  })

  // pay this recently created order
  payedOrder, _ := orderFactory.Pay(order.ID, &payments.OrderPayParams{
    Customer: customer.ID,
  })

  fmt.Println(payedOrder.ID)
}
