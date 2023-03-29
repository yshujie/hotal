package service

import "github.com/hotel/app/entity"

type customerService struct{}

// 表名
func (this customerService) table() string {
	return tableName("customer")
}

// GetCustomer 获取一个顾客
func (this customerService) GetCustomer(id int64) (*entity.Customer, error) {
	customer := &entity.Customer{}
	customer.Id = id
	if err := o.Read(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

// GetAllCustomer 获取所有的 customer
func (this customerService) GetAllCustomer() ([]entity.Customer, error) {
	return this.GetList(1, -1)
}

// GetList 获取 customer 列表
func (this customerService) GetList(page, pageSize int) ([]entity.Customer, error) {
	offset, limit := getPagination(page, pageSize)

	var list []entity.Customer
	_, err := o.QueryTable(this.table()).Offset(offset).Limit(limit).All(&list)
	return list, err
}

// GetTotal 获取 customer 总数
func (this customerService) GetTotal() (int64, error) {
	return o.QueryTable(this.table()).Count()
}

// AddCustomer 新增 customer
func (this customerService) AddCustomer(customer *entity.Customer) error {
	_, err := o.Insert(customer)
	return err
}

// UpdateCustomer 更新 customer 信息
func (this customerService) UpdateCustomer(customer *entity.Customer, fields ...string) error {
	_, err := o.Update(customer, fields...)
	return err
}

// DeleteCustomer 删除 customer
func (this customerService) DeleteCustomer(customerId int64) error {
	customer, err := this.GetCustomer(customerId)
	if err != nil {
		return err
	}

	o.Delete(customer)
	return nil
}
