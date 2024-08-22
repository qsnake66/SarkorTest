package SarkorTest

import "time"

type Product struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
