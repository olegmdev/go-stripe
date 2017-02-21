package payments

type Customer struct {
  ID string

  // TODO: all necessary customer fields should be there
}

type CustomerParams struct {
	Balance int64
	Source *SourceParams
	Desc, Email string
	BusinessVatID string
}

type SourceParams struct {
	Card *CardParams
}

type CardParams struct {
	Name, Number, Month, Year string
	Address1, Address2, City, State string
}

type CustomerFactory struct {
  IPaymentProvider
}

// Create new customer
func (self *CustomerFactory) New(params *CustomerParams) (*Customer, error) {
  return self.NewCustomer(params)
}

// Retrieve a customer by ID
func (self *CustomerFactory) Get(customerId string) (*Customer, error) {
  return self.GetCustomer(customerId)
}

// Retrieve all existing customers
func (self *CustomerFactory) List() ([]*Customer) {
  return self.GetCustomers()
}
