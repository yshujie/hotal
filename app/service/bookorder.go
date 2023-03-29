package service

import "github.com/hotel/app/entity"

type bookOrderService struct{}

// 表名
func (s *bookOrderService) table() string {
	return tableName("bookorder")
}

// GetBookOrder 获取 bookOrder 信息
func (s *bookOrderService) GetBookOrder(bookOrderId int64) (*entity.BookOrder, error) {
	bookOrder := &entity.BookOrder{}
	bookOrder.Id = bookOrderId

	if err := o.Read(bookOrder); err != nil {
		return nil, err
	}

	return bookOrder, nil
}

// GetAllBookOrder 获取所有的 bookOrder
func (s *bookOrderService) GetAllBookOrder() ([]entity.BookOrder, error) {
	return s.GetList(1, -1)
}

// GetList 获取 bookOrder 列表
func (s *bookOrderService) GetList(page, pageSize int) ([]entity.BookOrder, error) {
	offset, limit := getPagination(page, pageSize)

	var list []entity.BookOrder
	_, err := o.QueryTable(s.table()).Offset(offset).Limit(limit).All(&list)
	for i, _ := range list {
		o.LoadRelated(&list[i], "Room")
	}

	return list, err
}

// GetTotal 获取 bookOrder 总数量
func (s *bookOrderService) GetTotal() (int64, error) {
	return o.QueryTable(s.table()).Count()
}

// AddBookOrder 新增 bookOrder
func (s *bookOrderService) AddBookOrder(bookOrder *entity.BookOrder) error {
	bookOrder.Amount = float64(bookOrder.Daycnt) * bookOrder.Room.Price
	_, err := o.Insert(bookOrder)
	return err
}

// UpdateBookOrder 更新 bookOrder
func (s *bookOrderService) UpdateBookOrder(bookOrder *entity.BookOrder, fields ...string) error {
	_, err := o.Update(bookOrder, fields...)
	return err
}

// DeleteBookOrder 删除 bookOrder
func (s *bookOrderService) DeleteBookOrder(bookOrderId int64) error {
	bookOrder, err := s.GetBookOrder(bookOrderId)
	if err != nil {
		return err
	}

	o.Delete(bookOrder)
	return nil
}
