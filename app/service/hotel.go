package service

import "github.com/hotel/app/entity"

type hotelService struct{}

// 表名
func (s *hotelService) table() string {
	return tableName("hotel")
}

// GetHotel 获取一个 hotel 信息
func (s *hotelService) GetHotel(id int64) (*entity.Hotel, error) {
	hotal := &entity.Hotel{}
	hotal.Id = id
	if err := o.Read(hotal); err != nil {
		return nil, err
	}

	return hotal, nil
}

// GetAllHotel 获取所有的 hotel
func (s *hotelService) GetAllHotel() ([]entity.Hotel, error) {
	return s.GetList(1, -1)
}

// GetList 获取 hotel 列表
func (s *hotelService) GetList(page, pageSize int) ([]entity.Hotel, error) {
	var list []entity.Hotel
	offset := 0
	if pageSize == -1 {
		pageSize = 1000
	} else {
		offset = (page - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
	}

	_, err := o.QueryTable(s.table()).Offset(offset).Limit(pageSize).All(&list)
	return list, err
}

// GetListWithRoom 获取 hotel 列表 + room 信息
func (s *hotelService) GetListWithRoom(page, pageSize int) ([]entity.Hotel, error) {
	var list []entity.Hotel
	offset := 0
	if pageSize == -1 {
		pageSize = 1000
	} else {
		offset = (page - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
	}

	_, err := o.QueryTable(s.table()).Offset(offset).Limit(pageSize).All(&list)

	for i, _ := range list {
		o.LoadRelated(&list[i], "Rooms")
		(&list[i]).RoomCnt = len(list[i].Rooms)
	}

	return list, err
}

// GetTotal 获取项目总数
func (s *hotelService) GetTotal() (int64, error) {
	return o.QueryTable(s.table()).Count()
}

// AddHotel 新增 hotel
func (s *hotelService) AddHotel(hotel *entity.Hotel) error {
	_, err := o.Insert(hotel)
	return err
}

// UpdateHotel 更新 hotel
func (s *hotelService) UpdateHotel(hotel *entity.Hotel, fields ...string) error {
	_, err := o.Update(hotel, fields...)
	return err
}

// DeleteHotel 删除 hotel
func (s *hotelService) DeleteHotel(hotelId int64) error {
	hotel, err := s.GetHotel(hotelId)
	if err != nil {
		return err
	}
	o.Delete(hotel)
	return nil
}
