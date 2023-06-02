package model

type Item struct {
	Id           uint   `gorm:"column:id;primary_key" json:"id"`
	ItemImage    string `gorm:"column:item_image;" json:"item_image"`
	ItemName     string `gorm:"column:item_name;" json:"item_name"`
	ItemPrice    string `gorm:"column:item_price;" json:"item_price"`
	ItemDesc     string `gorm:"column:item_desc;" json:"item_desc"`
	ItemFullDesc string `gorm:"column:item_full_desc;" json:"item_full_desc"`
}
