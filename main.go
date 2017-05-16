package main

import (
	_ "go-crawler/routers"
	//"github.com/astaxie/beego"
	util "go-crawler/util"
	//"strconv"
	//"fmt"
	"net/url"
)

func main() {
	//beego.Run()

	/*client := new(util.HttpClient)
	client.Init()
	client.Step1()*/

	/*int64Time, _ := strconv.ParseInt("1018272", 10, 64)
	str:=strconv.FormatInt(int64Time,16)
	fmt.Println(str)
	fmt.Println(strconv.ParseInt("f89a1", 16, 64))*/

	values := url.Values{}
	values.Add("deviceId","A0000062FFCA19")
	util.Post("http://mip.m.womai.com/tinker/tinkerPatch.action",values)
}