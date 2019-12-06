package api

import (
	"strings"
	"gopkg.in/gomail.v2"
)

type EmailParam struct {
	//邮箱服务器地址
	ServerHost string 
	//邮箱服务器端口
	ServerPort int 
	//发件人邮箱地址
	FromEmail string
	//发件人邮箱密码
	FromPasswd string
	//Toers 接收者邮件
	Toers string
	//CCers 抄送者邮件
	CCers string
}

var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

func InitEmail(ep *EmailParam)  {
	toers := []string{}

	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPasswd = ep.FromPasswd

	m = gomail.NewMessage()

	if len(ep.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	

	// if len(ep.CCers) != 0 {
	// 	for _, tmp := range strings.Split(ep.CCers, ",") {
	// 		toers = append(toers, strings.TrimSpace(tmp))
	// 	}
	// 	m.SetHeader("Cc", toers...)
	// }

	m.SetAddressHeader("From",fromEmail,"行星咖啡")
	m.SetAddressHeader("To",toers[0],"2")
}

func SendEmail(subject, body string)  error{
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(serverHost, serverPort, fromEmail,fromPasswd)
	err := d.DialAndSend(m)
	print(err)
	if err != nil{
		panic(err)
	}
	return err

}