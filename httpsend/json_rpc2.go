package httpsend

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//By Bill
// Json-Rpc2.0 @filecoin Rpc
const (
	VERSION_INFO = "2.0"
)

/**
	@params url  请求地址
	@params msg  消息(json字符串)比如 `{"name":"Bill"}`
                         Ps:也可以使用  BuildData  方法构建数据
												@params params 泛型list参数   []interface{}{"123",123}
                                                @params method 调用方法       call

	@params header 携带的Http头   map[string]string{}{
						   "Content-Type" : "application/json",
                         }
	@return 返回字符串
*/
func SendJsonReq(uri string, data string, headers map[string]string) string {
	res, err := SendRpc(uri, []byte(data), headers)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}

func SendRpc(url string, msg []byte, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//用于filecoin
func BuildData(params interface{}, method string) string {
	_id, _ := strconv.Atoi(GetRandomSring(4))
	data := ToJsonString(map[string]interface{}{
		"jsonrpc": VERSION_INFO,
		"method":  method,
		"params":  params,
		"id":      _id,
	})
	return data
}

var (
	randSeek = int64(1)
	l        sync.Mutex
)

func GetRandomSring(num int, str ...string) string {
	s := "0123456789"
	if len(str) > 0 {
		s = str[0]
	}
	l := len(s)
	r := rand.New(rand.NewSource(getRandSeek()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(s[x : x+1])
	}
	return buf.String()
}

func GetRandomChars(num int) string {
	ss := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return GetRandomSring(num, ss)
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}

func ToJsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
