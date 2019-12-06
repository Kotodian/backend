package api

/*
	咖啡种类的各种功能
*/
import (
	"class/models/types"
	"github.com/astaxie/beego"
)

type V3Controller struct {
	beego.Controller
}

type typesMessage struct {
	Id 				int 			`json:"id"`
	Code			string 			`json:"code"`
	Name			string 			`json:"name"`
}

func (c *V3Controller) Post()	{
	path := c.Ctx.Request.URL.Path
	switch path{
		case "/api/v3/order":
			//调用方法
			break
	}
}

func (c *V3Controller) Get()  {
	path := c.Ctx.Request.URL.Path
	switch path{
	case "/api/v3/getType":
		//调用方法
		c.GetType()
		break
	}
}
func (c *V3Controller) GetType()  {
	rb := ResponseBody{}
	co,err := types.GetType()
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