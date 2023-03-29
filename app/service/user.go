package service

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/hotel/app/entity"
	"github.com/hotel/app/libs"
)

type userService struct{}

func (this *userService) table() string {
	return tableName("user")
}

// GetUser 获取用户信息
func (this *userService) GetUser(userId int64, getRoleInfo bool) (*entity.User, error) {
	user := &entity.User{}
	user.Id = userId

	err := o.Read(user)
	return user, err
}

// GetUserByName 根据名字查找用户
func (this userService) GetUserByName(userName string) (*entity.User, error) {
	user := &entity.User{}
	user.UserName = userName

	err := o.Read(user, "UserName")
	return user, err
}

// GetTotal 获取用户总数
func (this userService) GetTotal() (int64, error) {
	return o.QueryTable(this.table()).Count()
}

// GetAllUser 获取所有用户
func (this userService) GetAllUser(getRoleInfo bool) ([]entity.User, error) {
	return this.GetUserList(1, -1, true)
}

// GetUserList 获取用户列表
func (this userService) GetUserList(page, pageSize int, getRoleInfo bool) ([]entity.User, error) {
	offset, limit := getPagination(page, pageSize)

	var users []entity.User
	_, err := o.QueryTable(this.table()).OrderBy("id").Limit(limit, offset).All(&users)
	return users, err
}

// AddUser 添加用户
func (this userService) AddUser(userName, realName, email, mobile, password string, sex int) (*entity.User, error) {
	if exists, _ := this.GetUserByName(userName); exists.Id > 0 {
		return nil, errors.New("用户名已存在")
	}

	user := entity.NewUser()
	user.UserName = userName
	user.RealName = realName
	user.Sex = sex
	user.Email = email
	user.Mobile = mobile
	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = password
	beego.Trace("user ...", user)

	_, err := o.Insert(user)
	return user, err
}

// UpdateUser 更新用户信息
func (this userService) UpdateUser(user *entity.User, fields ...string) error {
	if len(fields) < 1 {
		return errors.New("更新字段不能为空")
	}

	_, err := o.Update(user, fields...)
	return err
}

// ModifyPassword 修改密码
func (this userService) ModifyPassword(userId int64, password string) error {
	user, err := this.GetUser(userId, false)
	if err != nil {
		return err
	}

	user.Salt = string(utils.RandomCreateBytes(10))
	user.Password = libs.Md5([]byte(password + user.Salt))
	_, err = o.Update(user, "Sale", "Password")
	return err
}

// DeleteUser 删除用户
func (this userService) DeleteUser(userId int64) error {
	if userId == 1 {
		return errors.New("不允许删除管理员账户")
	}

	user := &entity.User{}
	user.Id = userId
	_, err := o.Delete(user)
	return err
}
