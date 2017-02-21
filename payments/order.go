package payments

type Order struct {
  ID string

  // TODO: all necessary order fields should be there
}

type OrderParams struct {
	Items    []*OrderItemParams
	Shipping *ShippingParams
}

type OrderItemParams struct {
	Description string
	Parent      string
	Type        string
}

type ShippingParams struct {
	Name    string
	Address *AddressParams
}

type AddressParams struct {
	Line1      string
	City       string
	PostalCode string
	Country    string
}

type OrderPayParams struct {
	Customer string
}

type OrderFactory struct {
  IPaymentProvider
}

func (self *OrderFactory) New(items []*OrderItemParams, params *OrderParams) (*Order, error) {
  return self.NewOrder(items, params)
}

func (self *OrderFactory) Pay(id string, params *OrderPayParams) (*Order, error) {
  return self.PayOrder(id, params)
}
