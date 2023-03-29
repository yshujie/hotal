package service

import "github.com/hotel/app/entity"

type roomService struct{}

// 表名
func (s *roomService) table() string {
	return tableName("room")
}

// GetRoom 获取一个 room 信息
func (s *roomService) GetRoom(id int64) (*entity.Room, error) {
	room := &entity.Room{}
	room.Id = id
	if err := o.Read(room); err != nil {
		return nil, err
	}

	return room, nil
}

// GetAllRoom 获取所有的 room
func (s *roomService) GetAllRoom() ([]entity.Room, error) {
	return s.GetList(1, -1)
}

// GetList 获取 room 列表
func (s *roomService) GetList(page, pageSize int) ([]entity.Room, error) {
	var list []entity.Room

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
		o.LoadRelated(&list[i], "Hotel")
	}

	return list, err
}

// GetTotal 获取 room 总数
func (s *roomService) GetTotal() (int64, error) {
	return o.QueryTable(s.table()).Count()
}

// GetTotalOfHotel 获取 hotel 的 room 数量
func (s *roomService) GetTotalOfHotel(hotel *entity.Hotel) (int64, error) {
	return o.QueryTable(s.table()).Filter("hospitalid", hotel.Id).Count()
}

// AddRoom 新增 room
func (s *roomService) AddRoom(room *entity.Room) error {
	_, err := o.Insert(room)
	return err
}

// UpdateRoom 更新 room
func (s *roomService) UpdateRoom(room *entity.Room, fields ...string) error {
	_, err := o.Update(room, fields...)
	return err
}

// DeleteRoom 删除 room
func (s *roomService) DeleteRoom(roomId int64) error {
	room, err := s.GetRoom(roomId)
	if err != nil {
		return err
	}

	// 删除
	o.Delete(room)
	return nil
}
