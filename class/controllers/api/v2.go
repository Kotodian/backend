package api

/*
咖啡订单等功能
*/

import (
	"encoding/json"
	"class/models/coffee"
	"github.com/astaxie/beego"
)
type V2Controller struct {
	beego.Controller
}

type CoffeeMessage struct {
	CoffeeId 	string 		`json:"coffee_id"`
	Name 		string 		`json:"name"`
	Value		float64		`json:"value"`
	Img			string		`json:"img"`
	Type		string		`json:"type"`
	Des			string		`json:"des"`
}
type TypeMessage struct {
	Type		string		`json:"type"`
}
type CoffeeIdMessage struct {
	CoffeeId	string		`json:"coffee_id"`
}
func (c *V2Controller) Post()	{
	path := c.Ctx.Request.URL.Path
	switch path{
		case "/api/v2/order":
			//调用方法
			break
		case "/api/v2/getcoffeebytype":
			c.GetCoffeeType()
			break
		case "/api/v2/getcoffeebyid":
			c.GetCoffeeByid()
			break
	}
}

func (c *V2Controller) Get()  {
	path := c.Ctx.Request.URL.Path
	switch path{
	case "/api/v2/getallcoffee":
		//调用方法
		c.GetAllCoffee()
		break
	}
}

func (c *V2Controller) GetAllCoffee() {
	rb := ResponseBody{}
	co,err := coffee.GetAllCoffee()
	if err != nil {
		rb.Code = 404
		rb.Message = "数据库读取失败"
	}else{
		rb.Code = 200
		rb.Message = "数据库读取成功"
		rb.Value = co
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

//
func (c *V2Controller) GetCoffeeType()	{
	rb := ResponseBody{}
	body := c.Ctx.Input.RequestBody
	typ := TypeMessage{}
	e := json.Unmarshal(body,&typ)
	var coffees []*coffee.Coffee
	if e == nil{
		coffees = coffee.FindCoffeeByType(typ.Type)
		rb.Code = 200
		rb.Message = "数据库读取成功"
		rb.Value = coffees
	}else{
		rb.Code = 404
		rb.Message = "类型未传到后台"
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
func (c *V2Controller) GetCoffeeByid()	{
	rb := ResponseBody{}
	body := c.Ctx.Input.RequestBody
	coffeeId := CoffeeIdMessage{}
	e := json.Unmarshal(body,&coffeeId)
	if e != nil{
		rb.Code = 400
		rb.Message = "参数错误"
		c.Data["json"] = &rb
		c.ServeJSON()
	}else{
		Cof, e := coffee.FindCoffeeById(coffeeId.CoffeeId)
		if e == nil{
			rb.Code = 200
			rb.Message = "读取成功"
			rb.Value = Cof
			c.Data["json"] = &rb
			c.ServeJSON()
		}else{
			rb.Code = 404
			rb.Message = "读取失败"
			c.Data["json"] = &rb
			c.ServeJSON()
		}
	}
}