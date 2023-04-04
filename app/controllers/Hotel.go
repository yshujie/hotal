package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/hotel/app/entity"
	"github.com/hotel/app/service"
)

type HotelController struct {
	BaseController
}

func (c *HotelController) List() {
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}

	hotelList, err := service.HotelService.GetListWithRoom(page, c.pageSize)
	if err != nil {
		c.checkError(err)
	}

	beego.Trace("hotelList ...... ", hotelList)
	count, _ := service.HotelService.GetTotal()
	c.Data["count"] = count
	c.Data["hotelList"] = hotelList
	c.Data["pageBar"] = libs.NewPager()

	c.display()
}

func (c *HotelController) Add() {
	if c.isPost() {
		name := c.GetString("name")
		address := c.GetString("address")

		valid := validation.Validation{}
		valid.Required(name, "name").Message("请输入酒店名称")
		valid.Required(address, "address").Message("请输入酒店地址")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				c.showMsg(err.Message, MSG_ERR)
			}
		}

		hotel := entity.NewHotel()
		hotel.Name = name
		hotel.Address = address
		err := service.HotelService.AddHotel(hotel)
		c.checkError(err)
		c.redirect(beego.URLFor("HotelController.List"))
	}

	c.Data["pageTitle"] = "添加酒店"
	c.display()
}

func (c *HotelController) Modify() {
	hotelid, _ := c.GetInt64("hotelid")
	hotel, err := service.HotelService.GetHotel(hotelid)
	c.checkError(err)

	if c.isPost() {
		hotel.Name = c.GetString("name")
		hotel.Address = c.GetString("address")
		hotel.Remark = c.GetString("remark")
		hotel.Status, _ = c.GetInt8("status")
		service.HotelService.UpdateHotel(hotel)
		c.redirect(beego.URLFor("HotelController.List"))
	}

	c.Data["hotel"] = hotel
	c.Data["pageTitle"] = "修改酒店"
	c.display()
}
