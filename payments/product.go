package payments

import stripe "github.com/stripe/stripe-go"
import product "github.com/stripe/stripe-go/product"
import sku "github.com/stripe/stripe-go/sku"

// Create new product
func (conf *Config) CreateProduct(params *stripe.ProductParams) (*stripe.Product, error) {
  stripe.Key = conf.Key

  return product.New(params)
}

// Create new product with a SKU
func (conf *Config) CreateProductWithSKU(params *stripe.ProductParams, skuParams *stripe.SKUParams) (*stripe.Product, error) {
  stripe.Key = conf.Key

  target, err := product.New(params)
  if err != nil {
    return nil, err
  }

  skuParams.Product = target.ID
  if _, err := sku.New(skuParams); err != nil {
    return nil, err
  }

  return target, nil
}

// Retrieve a product by ID
func (conf *Config) GetProduct(productId string) (*stripe.Product, error) {
  stripe.Key = conf.Key

  return product.Get(productId)
}

// Retrieve all existing products
func (conf *Config) GetProducts() ([]*stripe.Product) {
  stripe.Key = conf.Key

  var products []*stripe.Product

  params := &stripe.ProductListParams{}
  iterator := product.List(params)

  for iterator.Next() {
    item := iterator.Product()
    products = append(products, item)
  }

  return products
}
