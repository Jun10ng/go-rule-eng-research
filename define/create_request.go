package define

type CreateOrderRequest struct {
	TotalAmount int64
	FinalPrice int64
	AdminFee int64
}

func (c *CreateOrderRequest)GetTotalAmount() int64 {
	return c.TotalAmount
}
func (c *CreateOrderRequest)GetFinalPrice() int64 {
	return c.FinalPrice
}
func (c *CreateOrderRequest)GetAdminFee() int64 {
	return c.AdminFee
}