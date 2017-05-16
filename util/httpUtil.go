package fetcher

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/juju/persistent-cookiejar"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"strconv"
	"encoding/base64"
	"encoding/json"
)

const (
	//userAgent      = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Mobile/14E304 MicroMessenger/6.5.7 NetType/WIFI Language/zh_CN"
	userAgent  = "MicroMessenger/6.5.7.1041 NetType/WIFI Language/zh_CN"
	connection = "keep-alive"
	pragma     = "no-cache"
	xhr        = "XMLHttpRequest"
	//acceptEncoding = "gzip"
	acceptEncoding = ""
	Referer        = "https://servicewechat.com/wx80f809371ae33eda/23/page-frame.html"
	charset        = "utf-8"

	platform = "4"

	open_src = "list"
	mobileno = "13661303427"
	lang     = "zh"
	citycode = "010"

	//Content-Length: 51
	//Connection: Keep-Alive
	//Accept-Encoding: gzip
)

var (
	logger     = Logger{Enabled: true}
	caCertPath = "D:/zhongliang/go/src/go-crawler/conf/mobike.cer"

	contentType = "application/x-www-form-urlencoded"
	referer     = "https://servicewechat.com/wx80f809371ae33eda/26/"
	accesstoken = "d67cbcb07c89bc187302ec168164c203"
	Host        = "cwx.mobike.com"
)

type HttpClient struct {
	Client *http.Client
}

type MyTransport struct {
	tr        http.RoundTripper
	BeforeReq func(req *http.Request)
	AfterReq  func(resp *http.Response, req *http.Request)
}

func MyNewTransportHttps(tr http.RoundTripper) *MyTransport {
	hasError := false

	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)

	//var cliCrt tls.Certificate //双向验证
	if err != nil {
		fmt.Println("ReadFile err:", err)
		hasError = true
	} else {
		pool.AppendCertsFromPEM(caCrt)
		/*
			//ca 双向验证
			cli, err := tls.LoadX509KeyPair("certs/cert_server/client.crt", "certs/cert_server/client.key")
			if err != nil {
				fmt.Println("Loadx509keypair err:", err)
				hasError = true
			} else {
				cliCrt = cli
			}*/
	}

	t := &MyTransport{}
	if tr == nil {
		if hasError {
			tr = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		} else {
			tr = &http.Transport{
				TLSClientConfig: &tls.Config{
					//InsecureSkipVerify: true
					RootCAs: pool,
					//ca双向验证
					//Certificates: []tls.Certificate{cliCrt},
				},
			}
		}
		//tr = http.DefaultTransport
	}
	t.tr = tr
	return t
}

func MyNewTransport(tr http.RoundTripper) *MyTransport {
	t := &MyTransport{}
	if tr == nil {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		//tr = http.DefaultTransport
	}
	t.tr = tr
	return t
}

func (t *MyTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t.BeforeReq(req)
	resp, err = t.tr.RoundTrip(req)
	if err != nil {
		return
	}
	t.AfterReq(resp, req)
	return
}

func (this *HttpClient) Init() {
	cookieJar, _ := cookiejar.New(nil)
	tr := MyNewTransport(nil)
	tr.AfterReq = func(resp *http.Response, req *http.Request) {
	}
	tr.BeforeReq = func(req *http.Request) {
	}

	if this.Client == nil {
		this.Client = &http.Client{
			Transport: tr,
			Jar:       cookieJar,
		}
	}
}

func (this *HttpClient) SaveCookie() {
	this.Client.Jar.(*cookiejar.Jar).Save()
}

func newHTTPHeaders(isXhr bool) http.Header {
	headers := make(http.Header)
	//headers.Set("Accept", "*/*")
	headers.Set("User-Agent", userAgent)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("referer", Referer)

	headers.Set("Host", Host)
	//headers.Set("Connection", connection)
	//headers.Set("Accept-Encoding", acceptEncoding)

	headers.Set("charset", charset)
	headers.Set("platform", platform)
	headers.Set("open_src", open_src)
	timestamp := time.Now().Unix()
	headers.Set("time", strconv.FormatInt(timestamp,10))
	headers.Set("mobileno", mobileno)
	headers.Set("lang", lang)
	headers.Set("citycode", citycode)

	headers.Set("accesstoken", accesstoken)

	if isXhr {
		headers.Set("X-Requested-With", xhr)
	}
	return headers
}

// Get 发起一个 GET 请求，自动处理 cookies
func (this *HttpClient) Get(url string, params string) (*http.Response, error) {
	logger.Info("GET %s", url)
	req, err := http.NewRequest("GET", url, strings.NewReader(params))
	if err != nil {
		logger.Error("NewRequest failed with URL: %s", url)
		return nil, err
	}

	req.Header = newHTTPHeaders(false)
	return this.Client.Do(req)
}

//
func values2string(values *url.Values) string {
	return values.Encode()
}

// Post 发起一个 POST 请求，自动处理 cookies
func (this *HttpClient) Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	logger.Info("POST %s, %s", url, bodyType)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	headers := newHTTPHeaders(false)
	headers.Set("Content-Type", bodyType)
	eption   := "f89a0"

	headers.Set("eption", eption)
	req.Header = headers
	return this.Client.Do(req)
}

func (this *HttpClient) Step1() bool {
	u := "https://mwx.mobike.com/mobike-api/rent/nearbyBikesInfo.do"
	values := url.Values{}
	//values = map[string]string{}
	values.Add("longitude", "116.429886")
	values.Add("latitude", "39.90645")
	values.Add("citycode", "010")
	values.Add("speed", "0")
	values.Add("errMsg", "getLocation:ok")
	values.Add("accuracy", "30")

	resp, err := this.Post(u,contentType ,strings.NewReader(values2string(&values)))
	if err != nil {
		logger.Error("登录失败：" + err.Error())
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("状态出错 StatusCode = %d", resp.StatusCode))
		return false
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败 %s", err.Error())
		return false
	}
	//this.SaveCookie()
	logger.Info("相应内容=%s", content)
	return true
}

func makeEption() {
	//1494569219805

	/*timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	str:=strconv.FormatInt(timestamp,16)
	fmt.Println(str)*/

	//#string到int64  f89a0
	int64Time, _ := strconv.ParseInt("1494569219", 10, 64)
	str:=strconv.FormatInt(int64Time,16)
	fmt.Println(str)
	fmt.Println(strconv.ParseInt("f89a0", 16, 64))
}

func Main() {
	//client := new(HttpClient)
	//client.Init()
	//client.Step1()
	values := make(map[string]string)
	values["deviceId"]="A0000062FFCA19"
	values["chnanle"]="2031"
	values["version"]="2.0.0"
	values["platform"]="android"
	Post("http://mip.m.womai.com/tinker/tinkerPatch.action",values)
}

//post数据
func Post(u string, body interface{}) (err error){
	client := new(HttpClient)
	client.Init()
	//values := url.Values{}
	values := map[string]interface{}{
		"common":"",
		"data":body,
	}
	b,err := json.Marshal(values)

	logger.Info(string(b))
	if err!=nil{
		logger.Error(err.Error())
		return err
	}
	data := base64Encode(b)
	v := url.Values{}
	v.Add("data",string(data))

	resp, err := client.Post(u,contentType ,strings.NewReader(v.Encode()))
	if err != nil {
		logger.Error("登录失败：" + err.Error())
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("状态出错 StatusCode = %d", resp.StatusCode))
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败 %s", err.Error())
		return err
	}
	client.SaveCookie()
	logger.Info("相应内容=%s", content)

	m:=make(map[string]string)
	e:=json.Unmarshal(content,&m)
	if e!=nil{
		panic(e)
	}
	b1,e := base64Decode([]byte(m["data"]))
	if e!=nil{
		panic(e)
	}
	logger.Info(string(b1))
	return err
}

//二进制转十六进制
func btox(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 16)
}

//十六进制转二进制
func xtob(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}
//用base64进行编码
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))

}

//用base64进行解码
func base64Decode(src []byte) ([]byte, error) {

	return base64.StdEncoding.DecodeString(string(src))

}