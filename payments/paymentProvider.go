package payments

type IPaymentProvider interface {
  NewCustomer(*CustomerParams) (*Customer, error)
  GetCustomer(string) (*Customer, error)
  GetCustomers() ([]*Customer)

  NewProduct(*ProductParams) (*Product, error)
  NewProductWithSKU(*ProductParams, *SKUParams) (*Product, error)
  GetProduct(string) (*Product, error)
  GetProducts() ([]*Product)

  NewOrder([]*OrderItemParams, *OrderParams) (*Order, error)
  PayOrder(string, *OrderPayParams) (*Order, error)
}
