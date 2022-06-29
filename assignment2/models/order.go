package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	ID           uint   `gorm:"primary_key"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Items        []Item `json:"items"`
}

func (ord *Order) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before order is created")
	if ord.OrderedAt == "" {
		ord.OrderedAt = time.Now().String()
	}

	errMsg := ""

	if len(ord.CustomerName) == 0 {
		errMsg = errMsg + "customer_name must be filled, "
	}

	if len(ord.Items) == 0 {
		errMsg = errMsg + " At least one item must be filled"
	}

	for i, _ := range ord.Items {
		if len(ord.Items[i].ItemCode) == 0 {
			errMsg = errMsg + " item_code must be filled,"
		}

		isDescEmpty := len(ord.Items[i].Description) == 0
		isQuantityEmpty := ord.Items[i].Quantity == 0
		chekCondition := isDescEmpty || isQuantityEmpty
		if chekCondition == true {
			errMsg = errMsg + " Item description or quantity must be filled."
		}
	}

	if errMsg != "" {
		err = errors.New(errMsg)
	}

	return
}

func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before order is update")

	errMsg := ""
	for index, _ := range order.Items {
		if order.Items[index].LineItemId != 0 {
			order.Items[index].OrderId = order.ID
			order.Items[index].ID = order.Items[index].LineItemId
		} else {
			errMsg = errMsg + "lineItemId item #" + strconv.Itoa(index+1) + " can't be empty. "
		}
	}

	if errMsg != "" {
		err = errors.New(errMsg)
	}

	return
}
