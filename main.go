package main

import "encoding/json"
import "fmt"
import "os"
import "./payments"

// TODO: still need to be incapsulated
import stripe "github.com/stripe/stripe-go"

func printEntity(entity interface{}) {
  res, _ := json.Marshal(entity)
  os.Stdout.Write(res)
}

func main() {
  fmt.Println("Running commands...")

  // TODO: gonfig should be used, put configuration in a file
  config := payments.Config{ Key: "sk_test_2lVxPRZeeKaeEncem5TlKsln" }
  client := payments.Payments(&config)

  // create new customer
  customer, _ := client.CreateCustomer(&stripe.CustomerParams{
		Balance:       10000,
		Desc:          "Test Customer",
		Email:         "a@b.com",
		BusinessVatID: "123456789",
    Source:        &stripe.SourceParams{
      Card:     &stripe.CardParams{
        Name:     "Stripe Tester",
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

  // create few products
  client.CreateProductWithSKU(&stripe.ProductParams{
    Name: "T-shirt",
  }, &stripe.SKUParams{
    Price: 150,
    Currency: "usd",
    Inventory: stripe.Inventory{Type: "infinite"},
  })

  client.CreateProductWithSKU(&stripe.ProductParams{
    Name: "Socks",
  }, &stripe.SKUParams{
    Price: 200,
    Currency: "usd",
    Inventory: stripe.Inventory{Type: "infinite"},
  })

  client.CreateProductWithSKU(&stripe.ProductParams{
    Name: "Hat",
  }, &stripe.SKUParams{
    Price: 600,
    Currency: "usd",
    Inventory: stripe.Inventory{Type: "infinite"},
  })

  // Get all existing products and parse SKUs
  products := client.GetProducts()
  var skuItems []*stripe.OrderItemParams
  for _, product := range products {
    skuItems = append(skuItems, &stripe.OrderItemParams{
      Type:   "sku",
      Parent: product.Skus.Values[0].ID,
      Description: "Some description",
    })
  }

  // create an order
  order, _ := client.CreateOrder(&stripe.OrderParams{
    Currency: "usd",
    Items: skuItems,
    Shipping: &stripe.ShippingParams{
      Name: "Noah Davis",
      Address: &stripe.AddressParams{
        Line1: "1234 Main Street",
        City: "San Francisco",
        Country: "US",
        PostalCode: "94111",
      },
    },
  })

  // pay this recently created order
  payedOrder, _ := client.PayOrder(order.ID, &stripe.OrderPayParams{
    Customer: customer.ID,
	})

  printEntity(payedOrder)
}
