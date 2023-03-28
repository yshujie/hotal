package entity

import "time"

type baseEntity struct {
	Id         int64
	Version    int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}
