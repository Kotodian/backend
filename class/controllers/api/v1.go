package api
/*
	用户登录 注册 等功能实现
*/
import (
	"github.com/astaxie/beego"
	"encoding/json"
	"class/models/user"
	"math/rand"
	"time"
)

const SessionUserKey = "heilian.SESSION_USERKEY"

type V1Controller struct {
	beego.Controller
}

type UserMessage struct {
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
	Image		string		`json:"image"`
}

func (c *V1Controller) Post() {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/api/v1/login":
		c.LoginPost()
		break
	case "/api/v1/register":
		c.RegisterPost()
		break
	case "/api/v1/emailsend":
		c.EmailSend()
		break
	case "/api/v1/phonesend":
		c.PhoneSend()
		break
	}
}

func (c *V1Controller) Get()  {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/api/v1/user":
		c.UserGet()
		break
	case "/api/v1/login-out":
		c.LoginOutGet()
		break
	}
}

func (c *V1Controller) LoginPost (){
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	e := json.Unmarshal(body,&um)
	rb := ResponseBody{}
	if e != nil{
		rb.Code = 400
		rb.Message = "参数错误"
		c.Data["json"] = &rb
		c.ServeJSON()
	}
	u,e := user.Login(um.Username,um.Password)
	if e == user.UsernameNotFind {
		rb.Code = 404
		rb.Message = "用户名不存在"
	} else if e == user.PasswordError{
		rb.Code = 404
		rb.Message = "密码错误"
	} else if e == nil {
		rb.Code = 200
		rb.Message = "登录成功"
		rb.Value = u.Image
		c.SetSession(SessionUserKey,&u)
	} else {
		rb.Code = 500
		rb.Message = e.Error()
	}
	c.Data["json"] = &rb
	print(u.Image)
	c.ServeJSON()
}

func (c *V1Controller) UserGet (){
	u := c.GetSession(SessionUserKey)
	rb := ResponseBody{}
	if u == nil {
		rb.Code = 404
		rb.Message = "当前没有登录用户"
	}else {
		rb.Code = 200
		rb.Message = "获取登录用户成功"
		rb.Value = &u
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
func (c *V1Controller) RegisterPost (){
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	e := json.Unmarshal(body,&um)
	rb := ResponseBody{}
	if e != nil {
		rb.Code = 400
		rb.Message = "参数错误"
		c.Data["json"] = &rb
		c.ServeJSON()
	}
	_,e = user.Register(um.Username,um.Password,um.Image)
	if e == user.UserExist {
		rb.Code = 404
		rb.Message = "用户已存在"
	}else if e == nil {
		rb.Code = 200
		rb.Message = "注册成功"
	}else {
		rb.Code = 500
		rb.Message = e.Error()
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

func (c *V1Controller) LoginOutGet (){
	c.DelSession(SessionUserKey)
	c.Data["json"] = &ResponseBody{Code:200,Message:"登出成功"}
	c.ServeJSON()
}

func (c *V1Controller) EmailSend (){
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	e := json.Unmarshal(body,&um)
	rb := ResponseBody{}
	if e != nil {
		rb.Code = 400
		rb.Message = "参数错误"
		c.Data["json"] = &rb
		c.ServeJSON()
	}
	
	serverHost := "smtp.qq.com"
	serverPort := 465
	fromEmail := "417350372@qq.com"
	fromPasswd := "itexnwidazxobiae"

	myToers := "asd254152@163.com"
	myCCers := ""
	identifyCode := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	subject := "注册账号验证码"
	emailBody := `验证码如下<br>
				<h3>`+string(identifyCode)+`</h3>`
	myEmail := &EmailParam {
		ServerHost: serverHost,
        ServerPort: serverPort,
        FromEmail:  fromEmail,
        FromPasswd: fromPasswd,
        Toers:      myToers,
        CCers:      myCCers,
	}
	InitEmail(myEmail)
	err := SendEmail(subject,emailBody)
	if err != nil{
		rb.Code = 400
		rb.Message = "发送失败"
		c.Data["json"] = &rb
		c.ServeJSON()
	}else{
		rb.Code = 200
		rb.Message = "发送成功"
		rb.Value = string(identifyCode)
		c.Data["json"] = &rb
		c.ServeJSON()
	}
}
func (c *V1Controller) PhoneSend()  {
	// body := c.Ctx.Input.RequestBody
	// um := UserMessage{}
	// e := json.Unmarshal(body,&um)
	rb := ResponseBody{}
	SMSSend()
	rb.Code = 200
	rb.Message = "发送成功"
	rb.Value = "2234"
	c.Data["json"] = &rb
	c.ServeJSON()
}