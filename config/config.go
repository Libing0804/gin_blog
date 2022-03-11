package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type TomlConfig struct {
	Viewer Viewer //这里的Viewer 是对应toml文件的viewer虽然大小写有区别 但是可以对应
	System SystemConfig
}

type Viewer struct {
	Title string
	Description string
	Logo string
	Navigation []string
	Bilibili string
	Avatar string
	UserName string
	UserDesc string
}
type SystemConfig struct {
	AppName string
	Version float32
	CurrentDir string
	CdnURL string
	//七牛云 存放数据
	QiniuAccessKey string
	QiniuSecretKey string
	//做评论用的组件
	Valine bool
	ValineAppid string
	ValineAppkey string
	ValineServerURL string

}

//从toml文件中直接加载固定的文件数据
var Cfg *TomlConfig
func init(){
	Cfg=new(TomlConfig)
	Cfg.System.AppName="gin成神之路"
	Cfg.System.Version=2.0
	//记录当前文件的目录
	Cfg.System.CurrentDir,_=os.Getwd()
	_,err:=toml.DecodeFile("config/config.toml",&Cfg)
	if err!=nil{
		log.Println(err)
		panic(err)
	}
}
