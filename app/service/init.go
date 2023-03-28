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
