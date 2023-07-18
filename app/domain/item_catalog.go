package domain

type ItemCatalog struct {
	ID          int     `gorm:"primaryKey"`
	Price       float64 `gorm:"price"`
	Name        string  `gorm:"name"`
	Description string  `gorm:"description"`
}

func (ItemCatalog) TableName() string {
	return "item_catalogs"
}