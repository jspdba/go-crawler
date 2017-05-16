package test

import (
	"testing"
	"net/url"
	util "go-crawler/util"
)

func Test_http_post(t *testing.T) {
	values := url.Values{}
	values.Add("deviceId","A0000062FFCA19")
	util.Post("http://mip.m.womai.com/tinker/tinkerPatch.action",values)
}