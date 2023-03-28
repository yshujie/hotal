package entity

import "time"

type Room struct {
	Id         int64
	Version    int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	Hotel      *Hotel    `json:"hotel" orm:"column(hotelid);ref(fk)"`
	No         string    `json:"no" orm:"size(64)"`
	Name       string    `json:"name" orm:"size(64)"`
	Price      float64   `json:"price" orm:"default(0)"`
	Status     int8      `json:"status" orm:"default(0)"`
	Remark     string    `json:"remark" orm:"size(256)"`
}

func NewRoom() *Room {
	return &Room{
		Id: getNextID(),
	}
}
