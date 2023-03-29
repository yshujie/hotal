package service

import "github.com/astaxie/beego/orm"

var (
	o orm.Ormer
)

// 获取真实表名
// todo：待实现
func tableName(name string) string {
	return name
}

// 获取分页参数
func getPagination(page, pageSize int) (offset, limit int) {
	if pageSize == -1 {
		offset = 0
		limit = 1000
	} else {
		offset = calcOffset(page, pageSize)
		limit = pageSize
	}

	return offset, limit
}

func calcOffset(page, pageSize int) (offset int) {
	if page <= 0 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	return offset
}
