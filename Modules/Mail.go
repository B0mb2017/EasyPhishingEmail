package Modules

type Mail struct {
	//邮件主题//
	Subject string
	//服务器IP
	SerHost string
	//发送的伪造发件人
	FromeAddr string
	//要调换的URL 前缀，先确认端口是否异常
	TargetUrl        string
	//DEBUG模式接受邮件的邮箱
	DebugEmail string
	//目标邮件组
	RecvLists []string
	//邮件发件人别名
	NickName string
}