package Plugin

import (
	"EasyPhishingEmail/Modules"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/farmerx/gorsa"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
//
//  ReadEMailLists
//  @Description: 将Excel的邮箱组添加到内存中，如果是DEBUG模式，
// 				  使用CONFIG_DEBUG_EMAIL接受邮件
//
func ReadEMailLists(mail *Modules.Mail,debug bool){
	var list []string
	if debug{
		list = append(list, mail.DebugEmail)
	} else {
		TranstableToSlice(mail)
	}
}
//
//  TranstableToSlice
//  @Description: 读取excel到slice
//
func TranstableToSlice(mail *Modules.Mail) {
	f, err := excelize.OpenFile("list.xlsx")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	//读取列信息
	sheets := f.GetSheetMap()
	if len(sheets) < 1 {
		log.Fatal("不能读取空表")
		os.Exit(0)
	}
	//获取第一个表的邮箱
	cols, err := f.Cols(sheets[1])
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	for cols.Next() {
		col, err := cols.Rows()
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(0)
		}
		for _, email := range col {
			if strings.Index(email, "@") != -1 {
				//将列表添加到Slice
				mail.RecvLists = append(mail.RecvLists, email)
			}
		}
	}
}
//
//  GeneralJIDText
//  @Description:  将邮件正文转为可追踪的邮件格式
//  @param body   模板正文 其中必须含有http://1.com,加密方式RSA取前48位
//  @param recv   将收件人邮箱进行加密，方便后期定位点击人
//  @return []byte
//
func GeneralJIDText(body []byte, recv string,mail *Modules.Mail) []byte {
	pubenctypt, err := gorsa.RSA.PubKeyENCTYPT([]byte(recv))
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	b64encode := base64.StdEncoding.EncodeToString(pubenctypt)
	var buffer bytes.Buffer
	buffer.WriteString(mail.TargetUrl)
	buffer.WriteString(b64encode[:48])
	fmt.Printf(recv + "\t" + b64encode[:48])
	//替换原有字符为目标字符，不限制替换数量
	newbody := bytes.Replace(body, []byte("http://1.com"), buffer.Bytes(), -1)
	return newbody
}
//
//  ReadBody
//  @Description: 读取钓鱼邮件正文模板
//  @return []byte
//
func ReadTpl() []byte {
	file, err := ioutil.ReadFile("tpl.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(file) <= 0 {
		log.Fatal("不能读取空模板!")
		os.Exit(0)
	}
	return file
}