package api

import (
	"fmt"
	sms "submail_go_sdk/submail/sms"
)

//发送短信
func SMSSend()  {
	config := make(map[string]string)
	config["appid"] = "43354"
	config["appkey"] = "2011cdde1108c65e5fc044e76816ef17"
	config["signType"] = "sha1"

	submail := sms.CreateSend(config)
	//测试
	submail.SetTo("13353305122")
	submail.SetContent("【行星咖啡】您的验证码是：2234，请在30分钟输入")

	send := submail.Send()
	fmt.Println("短信 Send 接口:",send)
}