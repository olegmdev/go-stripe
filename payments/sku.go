package payments

type SKU struct {
  ID string
}

type SKUParams struct {
	Price      int64
	Currency   string
	Inventory  Inventory
	Product    string
}

type SKUList struct {
  Values []*SKU
}

type Inventory struct {
	Type string

  // TODO: add Quantity and Value
}
