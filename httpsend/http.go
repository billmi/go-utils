package httpsend

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	CONN_TIME_OUT = time.Second * 2
)

/**
  Get
  @author Bill
*/
func Get(apiURL string, params url.Values) (resData string, e error) {
	var (
		Url *url.URL
		err error
	)
	Url, err = url.Parse(apiURL)
	if err != nil {
		return "", err
	}
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/**
Get with TimeOut
Code == 200 则返回,其他请打印错误
第三个参数设置超时  单位:秒 【0:则默认两秒】
*/
func SendGetWithTimeOut(apiUrl string, params url.Values, time_out int) (resData string, e error) {
	var (
		Url *url.URL
		err error
		b   []byte
	)
	Url, err = url.Parse(apiUrl)
	if err != nil {
		return "", err
	}
	Url.RawQuery = params.Encode()
	client := _httpClient(time_out)
	resp, err := client.Get(Url.String())
	if err != nil || resp == nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", nil
	}
	if resp.Body != nil {
		b, err = ioutil.ReadAll(resp.Body)
		return string(b), nil
	}
	return "", err
}

func _httpClient(time_out int) http.Client {
	var (
		_sec time.Duration
	)
	_sec = CONN_TIME_OUT
	if time_out > 0 {
		_sec = time.Second * time.Duration(time_out)
	}
	return http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, _sec)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(_sec))
				return conn, nil
			},
			ResponseHeaderTimeout: _sec,
		},
	}
}

/**
网络请求POST
*/
func Post(apiURL string, params url.Values) (resData string, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/**
Post with TimeOut
Code == 200 则返回,其他请打印错误
第三个参数设置超时  单位:秒 【0:则默认两秒】
*/
func SendPostWithTimeOut(apiUrl string, params url.Values, time_out int) (resData string, e error) {
	var (
		err error
		b   []byte
	)
	client := _httpClient(time_out)
	resp, err := client.PostForm(apiUrl, params)
	if err != nil || resp == nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", nil
	}
	if resp.Body != nil {
		b, err = ioutil.ReadAll(resp.Body)
		return string(b), nil
	}
	return "", err
}
