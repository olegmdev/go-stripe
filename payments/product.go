package payments

type Product struct {
  ID string
  Skus *SKUList

  // TODO: all necessary product fields should be there
}

type ProductParams struct {
  Name string
}

type ProductFactory struct {
  IPaymentProvider
}

// Create new product
func (self *ProductFactory) New(params *ProductParams) (*Product, error) {
  return self.NewProduct(params)
}

// Create new product with a SKU
func (self *ProductFactory) NewWithSKU(params *ProductParams, skuParams *SKUParams) (*Product, error) {
  return self.NewProductWithSKU(params, skuParams)
}

// Retrieve a product by ID
func (self *ProductFactory) Get(productId string) (*Product, error) {
  return self.GetProduct(productId)
}

// Retrieve all existing products
func (self *ProductFactory) List() ([]*Product) {
  return self.GetProducts()
}
