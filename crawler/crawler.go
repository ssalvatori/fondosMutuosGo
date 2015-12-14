package crawler

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/ssalvatori/fondosMutuosGo/fondosmutuos"
)

const sourceURL = "http://www.svs.cl/institucional/estadisticas/cfm_download.php"

/* headers
POST /institucional/estadisticas/cfm_download.php HTTP/1.1
Host: www.svs.cl
Connection: keep-alive
Content-Length: 76
Cache-Control: max-age=0
Origin: http://www.svs.cl
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.41 Safari/537.36
Content-Type: application/x-www-form-urlencoded
Referer: http://www.svs.cl/institucional/estadisticas/fondos_cartola_diaria.php
Accept-Encoding: gzip, deflate
Accept-Language: es,en;q=0.8,fr;q=0.6,it;q=0.4
Cookie: PHPSESSID=2794ff95a936bb3eb386376298c6cbde; _gat=1; _ga=GA1.2.408942410.1450031397

txt_inicio=11%2F12%2F2015&txt_termino=11%2F12%2F2015&ffmm=%25&captcha=SD8EE4
Name

*/

type Crawler struct {
	fmManager *fm.FmManager
}

func NewCrawler() *Crawler {

	c := new(Crawler)
	c.fmManager = fm.NewFmManager()

	return c
}

func (c *Crawler) Run() {

	urlData := url.Values{}
	urlData.Set("txt_inicio", "10/12/2015")
	urlData.Set("txt_termino", "10/12/2015")
	urlData.Set("ffmm", "%")

	c.downloadFile(sourceURL, urlData)

}

func (c *Crawler) downloadFile(url string, data url.Values) {
	resp, err := http.PostForm(url, data)

	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()

	fmt.Println("respon", resp.Request.URL)

	fmt.Println("%#v", resp.Request.Form)

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	if err != nil {
		fmt.Println(err)
	}

	fileName := randStr(10)
	filePath := "tmp/" + fileName

	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
	}
	defer output.Close()

	n, err := io.Copy(output, resp.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
	}

	fmt.Println(n, "bytes downloaded.")

}

func randStr(n int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
