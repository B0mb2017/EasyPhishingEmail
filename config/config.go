package config

import (
	"EasyPhishingEmail/Plugin"
	"fmt"
	"github.com/farmerx/gorsa"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//RSA加密公钥 2048bits
var Pubkey string
var Privatekey string
var processbar Plugin.Bar
//var DEBUG_MODE = false

func Init(){
	processbar.NewBarWithCustomAccuracy(0, 100, 1)
	Plugin.LoadLogo()
	processbar.Play(10)
	InitLogger()
	processbar.Play(15)
	InitEncoder()
	processbar.Finish()
	var input uint8
	fmt.Println("请确认以上信息，无需更改请输入Y or y")
	fmt.Scanln(&input)
	if input != 'Y' && input != 'y' {
		os.Exit(0)
	}
}
func InitLogger()  {
	file := "./" +"log"+ ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[Sender]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}
func InitEncoder(){
	baseDir,_:=os.Getwd()
	f, err := os.OpenFile(baseDir + "/public.pem",os.O_RDONLY,0755)
	defer f.Close()
	if err!=nil{
		log.Fatal("读取公钥失败")
		return
	}else{
		data,err := ioutil.ReadAll(f)
		if err!=nil {
			log.Fatal("读取公钥长度失败" ,err)
			return
		}
		Pubkey = string(data)
	}
	f, err = os.OpenFile(baseDir + "/private.pem",os.O_RDONLY,0755)
	if err!=nil{
		log.Fatal("读取私钥失败")
		return
	}else{
		data,err :=ioutil.ReadAll(f)
		if err!=nil {
			log.Fatal("读取私钥长度失败" ,err)
			return
		}
		Privatekey = string(data)
	}
	time.Sleep(100 * time.Millisecond)
	processbar.Play(45)
	if err := gorsa.RSA.SetPublicKey(Pubkey); err != nil {
		log.Fatal(`set public key :`, err)
	}
	if err := gorsa.RSA.SetPrivateKey(Privatekey); err != nil {
		log.Fatalln(`set private key :`, err)
	}
	time.Sleep(100 * time.Millisecond)
	processbar.Play(60)
}