package main

import (
	_ "go-crawler/routers"
	//"github.com/astaxie/beego"
	util "go-crawler/util"
	//"strconv"
	//"fmt"
	"fmt"
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

	/*
	values := make(map[string]string)
	values["deviceId"]="A0000062FFCA19"
	values["channle"]="2031"
	values["version"]="3.5.2"
	values["platform"]="android"

	//util.Post("http://localhost:8080/tinker/tinkerPatch.action",values)
	util.Post("http://mip.m.womai.com/tinker/tinkerPatch.action",values)
	*/
	//util.Down("https://github-cloud.s3.amazonaws.com/releases/11892946/63209fb0-6b74-11e6-8a0a-fbac7b27d3c1.exe?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAISTNZFOVBIJMK3TQ%2F20170516%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20170516T093039Z&X-Amz-Expires=300&X-Amz-Signature=07b59a08d1958d167aff1e4f121e85a04720bde6deb36e3600d676a392f05812&X-Amz-SignedHeaders=host&actor_id=0&response-content-disposition=attachment%3B%20filename%3Dredis-desktop-manager-0.8.8.384.exe&response-content-type=application%2Foctet-stream")
	//2017-05-11 15:54:09,104 INFO  [1494489249102] com.womai.m.mip.channel.hotFix.AppPatchController.decryptRequestData - Decrypted request data:{"common":{"cityCode":"31000","mid":"0","userSession":"","userId":"","test1":""},"data":{"platform":"android","deviceId":"000000000000000","errorResult":"UpgradePatch tryPatch:new patch recover, try patch dex failed","patchVersion":null,"brandModel":"Google Nexus 6 - 5.0.0 - API 21 - 1440x2560generic","sysRelease":"5.0","channle":"2024","success":"false"}}

	/*
	values := make(map[string]string)
	values["deviceId"]="A1111162FFCA19"
	values["channle"]="2031"
	values["version"]="3.5.2"
	values["platform"]="android"
	values["errorResult"]="UpgradePatch tryPatch:new patch recover, try patch dex failed"
	values["patchVersion"]=""
	values["brandModel"]="Google Nexus 6 - 5.0.0 - API 21 - 1440x2560generic"
	values["sysRelease"]="5.0"
	values["success"]="false"

	//util.Post("http://localhost:8080/tinker/tinkerPatch.action",values)
	util.Post("http://mip.m.womai.com/tinker/patchDownLog.action",values)
	*/


	data := "eyJjb21tb24iOnsiY2l0eUNvZGUiOiIzMTAwMCIsIm1pZCI6IjAiLCJ1c2VyU2Vzc2lvbiI6IiIsInVzZXJJZCI6IiIsInRlc3QxIjoiIn0sImRhdGEiOnsicGxhdGZvcm0iOiJhbmRyb2lkIiwiY2hhbm5sZSI6IjMwMDciLCJkZXZpY2VJZCI6Ijg2ODAyOTAyMDEzMzU4MyIsInZlcnNpb24iOiIzLjUuMiJ9fQ=="
	b1,e :=util.Base642String(data)

	if e!=nil{
		panic(e)
	}
	fmt.Print(string(b1))
}