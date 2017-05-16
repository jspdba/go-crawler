package main

import (
	_ "go-crawler/routers"
	//"github.com/astaxie/beego"
	util "go-crawler/util"
	//"strconv"
	//"fmt"
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

	values := make(map[string]string)
	values["deviceId"]="A0000062FFCA19"
	values["channle"]="2031"
	values["version"]="2.0.0"
	values["platform"]="android"

	util.Post("http://localhost:8080/tinker/tinkerPatch.action",values)
	//util.Down("https://github-cloud.s3.amazonaws.com/releases/11892946/63209fb0-6b74-11e6-8a0a-fbac7b27d3c1.exe?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAISTNZFOVBIJMK3TQ%2F20170516%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20170516T093039Z&X-Amz-Expires=300&X-Amz-Signature=07b59a08d1958d167aff1e4f121e85a04720bde6deb36e3600d676a392f05812&X-Amz-SignedHeaders=host&actor_id=0&response-content-disposition=attachment%3B%20filename%3Dredis-desktop-manager-0.8.8.384.exe&response-content-type=application%2Foctet-stream")
}