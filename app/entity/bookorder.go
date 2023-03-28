package entity

import "time"

type BookOrder struct {
	Id         int64
	Version    int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	Room       *Room     `json:"room" orm:"column(roomid);ref(fk)"`
	Customer   *User     `json:"customer" orm:"column(customerid);ref(fk)"`
	Daycnt     int       `json:"daycnt" orm:"default(0)"`
	Amount     float64   `json:"amount" orm:"default(0)"`
}

func NewBookOrder() *BookOrder {
	return &BookOrder{
		Id: getNextID(),
	}
}
