package main

import (
	"EasyPhishingEmail/Modules"
	"EasyPhishingEmail/Plugin"
	"EasyPhishingEmail/config"
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"log"
	"os"
)
var Mail = Modules.Mail{
	Subject: "密码强度过低",
	SerHost: "127.0.0.1",
	FromeAddr: "test@domain.com",
	TargetUrl: "http://test.domain.com/?JSESSIONID=",
	DebugEmail: "test@domain.com",
	NickName: "测试管理员",
	RecvLists: []string{"test@domain.com"},
}
func SendMail(){
	var MailHelper *gomail.Message
	MailHelper = gomail.NewMessage()
	mailtpl := Plugin.ReadTpl()
	Plugin.ReadEMailLists(&Mail,true)
	MailHelper.SetAddressHeader("From", Mail.FromeAddr, Mail.NickName)
	MailHelper.SetHeader("To", Mail.RecvLists...)
	MailHelper.SetHeader("Subject", Mail.Subject)
	delivery := gomail.NewDialer(Mail.SerHost, 465, "yourusername", "yourpassword")
	delivery.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//发送错误次数
	errnum := 0
	for i := 0; i < len(Mail.RecvLists); i++ {
		body := Plugin.GeneralJIDText(mailtpl, Mail.RecvLists[i],&Mail)
		MailHelper.SetBody("text/html", string(body))
		if err := delivery.DialAndSend(MailHelper); err != nil {
			//连续发送三封邮件都失败了 那可能被对面网关ban了，终止发送，人工排查
			log.Fatal(err)
			errnum++
			if errnum >= 3 {
				os.Exit(0)
			}
		}
	}
}

func main() {
	config.Init()  //打印首屏Logo,初始化配置信息
	SendMail()
}
