package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hotel/app/entity"
	"github.com/hotel/app/libs"
	"github.com/hotel/app/service"
)

type BookOrderController struct {
	BaseController
}

func (c *BookOrderController) List() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	bookOrderList, err := service.BookOrderService.GetList(page, c.pageSize)
	c.checkError(err)

	count, _ := service.BookOrderService.GetTotal()
	c.Data["count"] = count
	c.Data["bookOrderList"] = bookOrderList
	c.Data["pageBar"] = libs.NewPager(page, int(count), c.pageSize, beego.URLFor("BookOrderController.List"), true).ToString()

	c.display()
}

func (c *BookOrderController) Book() {
	roomId, _ := c.GetInt64("roomid")
	room, err := service.RoomService.GetRoom(roomId)
	c.checkError(err)

	if c.isPost() {
		dayCnt, _ := c.GetInt("daycnt")

		user := c.auth.GetUser()
		bookOrder := entity.NewBookOrder()
		bookOrder.Room = room
		bookOrder.Daycnt = dayCnt
		bookOrder.Customer = user

		err = service.BookOrderService.AddBookOrder(bookOrder)
		c.checkError(err)
	}

	c.Data["room"] = room

	c.display()
}