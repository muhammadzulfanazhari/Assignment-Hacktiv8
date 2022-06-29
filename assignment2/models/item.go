package models

type Item struct {
	ID          uint   `gorm:"primary_key" json:"item_id"`
	LineItemId  uint   `json:"line_item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint   `json:"order_id"`
}
