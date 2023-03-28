package entity

import "time"

type Hotel struct {
	Id          int64
	Version     int
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	Name        string    `json:"name" orm:"size:(64)"`
	Address     string    `json:"address" orm:"size(256)"`
	AuditStatus int8      `json:"auditstatus" orm:"column(auditstatus);default(0)"`
	AuditRemark string    `json:"auditremark" orm:"column(auditremark);size(256);type(text)"`
	Status      int8      `json:"status" orm:"default(0)"`
	Remark      string    `json:"remark" orm:"size(256);type(text)"`
	Rooms       []*Room   `json:"room" orm:"reverse(many)"`
	RoomCnt     int       `json:"roomCnt" orm:"-"`
}

func NewHotel() *Hotel {
	return &Hotel{
		Id: getNextID(),
	}
}

// TableName 自定义表明
func (this *Hotel) TableName() string {
	return "hotel"
}
