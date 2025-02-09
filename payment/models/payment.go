package models

type Payment struct {
	ID      uint `gorm:"primaryKey"`
	IsPayed bool
	OrderId uint
}
