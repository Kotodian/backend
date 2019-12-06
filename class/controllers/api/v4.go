package api

import (
	"encoding/json"
	"class/models/address"
	"github.com/astaxie/beego"
)

type V4Controller struct {
	beego.Controller
}

type UserAddressMessage struct {
	Id 				int 			`json:"id"`
	Address			string 			`json:"address"`
	Sex				string			`json:"sex"`
	Username		string 			`json:"username"`
	Nickname		string 			`json:"nickname"`
	Phone			string			`json:"phone"`
	Isdefault		string			`json:"isdefault"`
}

func (c *V4Controller) Post() {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/api/v4/getAddress":
		c.GetAddress();
		break
	case "/api/v4/saveAddress":
		c.SaveAddress()
		break
	case "/api/v4/getAddressbyid":
		c.GetAddressById()
		break
	}

}
func (c *V4Controller) GetAddress(){
	body := c.Ctx.Input.RequestBody
	rb := ResponseBody{}
	var userAddress []*address.UserAddress
	user := UserAddressMessage{}
	e := json.Unmarshal(body,&user)
	if e == nil{
		userAddress = address.GetAddress(user.Username)
		rb.Code = 200
		rb.Message = "读取成功"
		rb.Value = userAddress
	}else{
		rb.Code = 400
		rb.Message = "未获取到用户名"
		rb.Value = e.Error
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
func (c *V4Controller) SaveAddress(){
	body := c.Ctx.Input.RequestBody
	udm := UserAddressMessage{}
	rb := ResponseBody{}
	e := json.Unmarshal(body,&udm)
	print(e)
	if e == nil{
		err := address.SaveAddress(udm.Address,udm.Username,udm.Nickname,udm.Phone,udm.Sex,udm.Isdefault)
		if err == nil{
			rb.Code = 200
			rb.Message = "插入成功"
		}else{
			rb.Code = 404
			rb.Message = "插入失败"
		}
	}else{
		rb.Code = 404
		rb.Message = "json解析失败"
		rb.Value = e
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

func (c *V4Controller) GetAddressById()  {
	body := c.Ctx.Input.RequestBody
	udm := UserAddressMessage{}
	rb := ResponseBody{}
	e := json.Unmarshal(body,&udm)
	if(e == nil){
		adre := address.UserAddress{}
		adre = address.GetAddressById(udm.Id)
		rb.Code = 200
		rb.Message = "读取成功"
		rb.Value = adre
	}else{
		rb.Code = 404
		rb.Message = "无结果"
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
