package models

type Cart struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint
	Item   []Item
}

type Item struct {
	ID        uint `gorm:"primaryKey"`
	ProductId uint
	Quantity  uint
	CartID    uint
}
