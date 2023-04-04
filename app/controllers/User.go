package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/hotel/app/service"
	"regexp"
)

type UserController struct {
	BaseController
}

func (c *UserController) List() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}

	userList, err := service.UserService.GetAllUser(false)
	c.checkError(err)

	c.Data["userList"] = userList
	c.display()
}

func (c *UserController) Add() {
	if c.isPost() {
		valid := validation.Validation{}

		username := c.GetString("username")
		realname := c.GetString("realname")
		email := c.GetString("email")
		mobile := c.GetString("mobile")
		sex, _ := c.GetInt("sex")
		password1 := c.GetString("password1")
		password2 := c.GetString("password2")

		valid.Required(username, "username").Message("请输入登录名")
		valid.Required(realname, "realname").Message("请输入真实姓名")
		valid.Required(email, "email").Message("请输入Email")
		valid.Required(password1, "password1").Message("请输入密码")
		valid.Required(password2, "password2").Message("请输入确认密码")
		valid.MinSize(password1, 6, "password1").Message("密码长度不能小于6个字符")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				c.showMsg(err.Message, MSG_ERR)
			}
		}

		_, err := service.UserService.AddUser(username, realname, email, mobile, password1, sex)
		c.checkError(err)

		c.redirect(beego.URLFor("UserController.List"))
	}

	c.Data["pageTitle"] = "添加账号"
	c.display()
}

func (c *UserController) Modify() {
	userId, _ := c.GetInt64("userid")
	user, err := service.UserService.GetUser(userId)
	c.checkError(err)

	if c.isPost() {
		valid := validation.Validation{}

		realname := c.GetString("realname")
		email := c.GetString("email")
		mobile := c.GetString("mobile")
		sex, _ := c.GetInt("sex")
		status, _ := c.GetInt("status")
		password1 := c.GetString("password1")
		password2 := c.GetString("password2")

		valid.Required(realname, "realname").Message("请输入真实姓名")
		valid.Required(email, "email").Message("请输入Email")
		valid.Email(email, "email").Message("Email无效")

		if password1 != "" {
			valid.Required(password1, "password1").Message("请输入密码")
			valid.Required(password2, "password2").Message("请输入确认密码")
			valid.MinSize(password1, 6, "password1").Message("密码长度不能小于6个字符")
			valid.Match(password1, regexp.MustCompile(`^`+regexp.QuoteMeta(password2)+`$`), "password2").Message("两次输入的密码不一致")
		}

		if valid.HasErrors() {
			for _, err := range valid.Errors {
				c.showMsg(err.Message, MSG_ERR)
			}
		}

		user.RealName = realname
		user.Sex = sex
		user.Status = status
		user.Email = email
		user.Mobile = mobile

		service.UserService.UpdateUser(user, "RealName", "Sex", "Status", "Email", "Mobile")
		c.redirect(beego.URLFor("UserController.List"))
	}

	c.Data["pageTitle"] = "修改账号"
	c.Data["user"] = user

	c.display()
}