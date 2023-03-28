package entity

import "time"

type Customer struct {
	Id         int64
	Version    int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	Name       string    `json:"name" orm:"size(64)"`
	Mobile     string    `json:"mobile" orm:"size(64)"`
}

func NewCustomer() *Customer {
	return &Customer{
		Id: getNextID(),
	}
}
