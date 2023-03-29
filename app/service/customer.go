package service

import "github.com/hotel/app/entity"

type customerService struct{}

// 表名
func (s *customerService) table() string {
	return tableName("customer")
}

// GetCustomer 获取一个顾客
func (s *customerService) GetCustomer(id int64) (*entity.Customer, error) {
	customer := &entity.Customer{}
	customer.Id = id
	if err := o.Read(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

// GetAllCustomer 获取所有的 customer
func (s *customerService) GetAllCustomer() ([]entity.Customer, error) {
	return s.GetList(1, -1)
}

// GetList 获取 customer 列表
func (s *customerService) GetList(page, pageSize int) ([]entity.Customer, error) {
	offset, limit := getPagination(page, pageSize)

	var list []entity.Customer
	_, err := o.QueryTable(s.table()).Offset(offset).Limit(limit).All(&list)
	return list, err
}

// GetTotal 获取 customer 总数
func (s *customerService) GetTotal() (int64, error) {
	return o.QueryTable(s.table()).Count()
}

// AddCustomer 新增 customer
func (s *customerService) AddCustomer(customer *entity.Customer) error {
	_, err := o.Insert(customer)
	return err
}

// UpdateCustomer 更新 customer 信息
func (s *customerService) UpdateCustomer(customer *entity.Customer, fields ...string) error {
	_, err := o.Update(customer, fields...)
	return err
}

// DeleteCustomer 删除 customer
func (s *customerService) DeleteCustomer(customerId int64) error {
	customer, err := s.GetCustomer(customerId)
	if err != nil {
		return err
	}

	o.Delete(customer)
	return nil
}
