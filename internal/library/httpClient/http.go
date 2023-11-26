package httpClient

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"gin/internal/global"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Timeout = 5
)

// get
func HttpGet(url string) ([]byte, error) {
	fmt.Println("--------url---------")
	fmt.Println(url)
	resp, err := http.Get(url)
	fmt.Println("--------1---------")
	fmt.Println(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("--------2---------")
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	return body, nil
}

func HttpGetWithHeader(httpUrl string, header map[string]string) string {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
			DialContext:           (&net.Dialer{Timeout: time.Second * 2}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(Timeout) * time.Second,
	}
	//读取Api数据
	req, err := http.NewRequest("GET", httpUrl, nil)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HttpGetWithHeaderProxy(httpUrl string, header map[string]string, proxy string) string {
	//proxy :=  url.Parse("http://12.23.16.11:1234")
	//transport := &http.Transport{Proxy: proxy}
	//c := &http.Client{Transport: transport}
	uri := url.URL{}
	proxyUrl, _ := uri.Parse(proxy)
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
			DialContext:           (&net.Dialer{Timeout: time.Second * 2}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
			Proxy: http.ProxyURL(proxyUrl),
		},
		Timeout: time.Duration(Timeout) * time.Second,
	}

	//读取Api数据
	req, err := http.NewRequest("GET", httpUrl, nil)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	//fmt.Println(1111)
	resp, err := client.Do(req)
	//defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// post
func HttpPost(url string, requestBody interface{}) ([]byte, error) {
	// 将请求数据转换为 JSON 字节
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http err encoding JSON:%s", err.Error()))
	}

	// 创建 POST 请求
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http response err:%s", err.Error()))
	}
	defer response.Body.Close()

	// 读取响应的内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http io.ReadAll err:%s", err.Error()))
	}

	return body, nil
}

func HttpPostWithHeader(http_url string, data map[string]string, header map[string]string) string {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
			DialContext:           (&net.Dialer{Timeout: time.Second * 2}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(Timeout) * time.Second,
	}
	//body
	urlValues := url.Values{}
	for k, v := range data {
		urlValues.Add(k, v)
	}
	req, err := http.NewRequest("POST", http_url, strings.NewReader(urlValues.Encode()))
	if err != nil {
		panic(err)
	}
	//header
	//	header["content-type"]="application/x-www-form-urlencoded"
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	//获取cookie后并转化成string
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HttpSend(reqUrl string, method string, data map[string]string, header map[string]string, timeOut time.Duration) ([]byte, error) {
	reqTimeOut := time.Duration(global.Cfg.Http.Timeout) * time.Second
	if timeOut != 0 {
		reqTimeOut = timeOut
	}
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 3,
			DialContext:           (&net.Dialer{Timeout: time.Second * 2}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
		Timeout: reqTimeOut,
	}
	//body
	urlValues := url.Values{}
	for k, v := range data {
		urlValues.Add(k, v)
	}
	req, err := http.NewRequest(method, reqUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	//获取cookie后并转化成string
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

func HttpStreamSend(reqUrl string, method string, data map[string]string, header map[string]string, timeOut time.Duration, ch chan []byte) {
	reqTimeOut := time.Duration(global.Cfg.Http.StreamTimeout) * time.Second
	if timeOut != 0 {
		reqTimeOut = timeOut
	}
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 3,
			DialContext:           (&net.Dialer{Timeout: time.Second * 2}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
		Timeout: reqTimeOut,
	}
	//body
	urlValues := url.Values{}
	for k, v := range data {
		urlValues.Add(k, v)
	}
	req, err := http.NewRequest(method, reqUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		ch <- []byte(err.Error())
		return
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	//获取cookie后并转化成string
	receive := bufio.NewReader(resp.Body)
	for {
		content, err := receive.ReadBytes('\n')
		if err != nil {
			break
		}
		ch <- content
	}

}