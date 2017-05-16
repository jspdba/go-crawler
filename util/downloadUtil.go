package fetcher

import (
	"net/http"
	"os"
	"io"
	"crypto/tls"
)

func Down(u string) {

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get(u)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("a.exe")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}
func main() {
	Down("https://github-cloud.s3.amazonaws.com/releases/11892946/63209fb0-6b74-11e6-8a0a-fbac7b27d3c1.exe?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAISTNZFOVBIJMK3TQ%2F20170516%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20170516T093039Z&X-Amz-Expires=300&X-Amz-Signature=07b59a08d1958d167aff1e4f121e85a04720bde6deb36e3600d676a392f05812&X-Amz-SignedHeaders=host&actor_id=0&response-content-disposition=attachment%3B%20filename%3Dredis-desktop-manager-0.8.8.384.exe&response-content-type=application%2Foctet-stream")
}
