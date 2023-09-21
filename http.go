package gotools

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetResponse(urls string) ([]byte, error) {
	maxRetry := 3
	i := 0
	ua := GetUA()
	for {
		if i > maxRetry {
			return nil, errors.New("获取内容为空")
		}

		request, _ := http.NewRequest("GET", urls, nil)
		//随机返回User-Agent 信息
		request.Header.Set("User-Agent", ua.getPcAgent())
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Set("Connection", "keep-alive")
		//设置超时时间
		timeout := time.Duration(20 * time.Second)
		client := &http.Client{
			Timeout: timeout,
		}
		response, err := client.Do(request)
		if response != nil {
			defer response.Body.Close()
		}
		if err != nil || response.StatusCode != 200 {
			i++
			continue
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			i++
			continue
		}
		return body, nil
	}

}

func HttpPost(urls string, data map[string]string) (string, error) {
	var c *http.Client = &http.Client{

		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					fmt.Println("dail timeout", err)
					return nil, err
				}
				return c, nil

			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	var clusterInfo = url.Values{}
	for k, v := range data {
		clusterInfo.Add(k, v)
	}
	params := clusterInfo.Encode()
	resp, err := c.Post(urls, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPostJson(url string, data string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", err
	}

	if req == nil {
		return "", err
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	if resp == nil {
		return "", err
	}
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Dingding(msg string) {
	message := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	resp, _ := HttpPostJson("https://oapi.dingtalk.com/robot/send?access_token=31d0eb905800f9b253f0c926f43d12f1b0cf90fd9126a7ca720ad65733cedac8", message)
	fmt.Println(resp)
}

// 获得html内容
func GetHtml(url string) ([]byte, error) {
	var logInfo string
	response, err := GetResponse(url)
	if err != nil {
		logInfo = fmt.Sprintf("http 请求失败，url:%v,err:%v", url, err)
		return nil, errors.New(logInfo)
	}

	return response, nil
}
