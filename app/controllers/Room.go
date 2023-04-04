package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/hotel/app/entity"
	"github.com/hotel/app/libs"
	"github.com/hotel/app/service"
)

type RoomController struct {
	BaseController
}

func (c *RoomController) List() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}

	hotelid, _ := c.GetInt64("hotalid")
	hotel, err := service.HotelService.GetHotel(hotelid)
	c.checkError(err)

	roomList, err := service.RoomService.GetListOfHotel(hotel, page, c.pageSize)
	if err != nil {
		c.checkError(err)
	}

	count, _ := service.RoomService.GetTotal()

	c.Data["hotel"] = hotel
	c.Data["count"] = count
	c.Data["roomList"] = roomList
	c.Data["pageBar"] = libs.NewPager(page, int(count), c.pageSize, beego.URLFor("RoomController.List"), true).ToString()

	c.display()
}

func (c *RoomController) Add() {
	hotelid, _ := c.GetInt64("hotelid")
	hotel, err := service.HotelService.GetHotel(hotelid)
	c.checkError(err)

	if c.isPost() {
		name := c.GetString("name")
		roomno := c.GetString("roomno")
		price, _ := c.GetFloat("price")
		remark := c.GetString("remark")

		valid := validation.Validation{}
		valid.Required(hotelid, "hotelid").Message("请输入酒店ID")
		valid.Required(name, "name").Message("请输入房间名称")
		valid.Required(roomno, "roomno").Message("请输入房号")
		valid.Required(price, "price").Message("请输入房间价格")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				c.showMsg(err.Message, MSG_ERR)
			}
		}

		room := entity.NewRoom()
		room.Hotel = hotel
		room.Name = name
		room.No = roomno
		room.Price = price
		room.Remark = remark
		c.checkError(service.RoomService.AddRoom(room))
		c.redirect(beego.URLFor("RoomController.List", "hotelid", hotel.Id))
	}

	c.Data["hotel"] = hotel
	c.Data["pageTitle"] = "添加房间"

	c.display()
}

func (c *RoomController) Modify() {
	roomId, _ := c.GetInt64("roomid")
	room, err := service.RoomService.GetRoom(roomId)
	c.checkError(err)

	if c.isPost() {
		room.Name = c.GetString("name")
		room.No = c.GetString("roomno")
		room.Price, err = c.GetFloat("price")
		room.Remark = c.GetString("remark")
		room.Status, _ = c.GetInt8("status")
		service.RoomService.UpdateRoom(room)
		c.redirect(beego.URLFor("RoomController.List", "hotelid", room.Hotel.Id))
	}

	c.Data["room"] = room
	c.Data["pageTitle"] = "修改酒店"

	c.display()
}