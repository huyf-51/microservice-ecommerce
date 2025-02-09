package types

type Cart struct {
	UserId uint `form:"userId"`
	Item   Item `form:"item"`
}

type Item struct {
	ProductId uint `form:"productId"`
	Quantity  uint `form:"quantity"`
}
