package entity

import "time"

type User struct {
	Id         int64
	Version    int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	UserName   string    `json:"username" orm:"unique;size(20)"`
	RealName   string    `json:"realname" orm:"size(32)"`
	Password   string    `json:"password" orm:"size(32)"`
	Salt       string    `json:"salt" orm:"size(10)"`
	Sex        int       `json:"sex" orm:"default(1)"`
	Email      string    `json:"email" orm:"size(50)"`
	Mobile     string    `json:"mobile" orm:"size(11)"`
	LastLogin  time.Time `json:"lastlogin" orm:"null;type(datetime)"`
	LastIp     string    `json:"lastip" orm:"size(15)"`
	Status     int       `json:"status" orm:"default(0)"`
}

func NewUser() *User {
	return &User{
		Id: getNextID(),
	}
}
