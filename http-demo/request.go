package http_demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get() {
	// 设置 InfluxDB 的地址和端口
	influxdbURL := "http://10.20.121.247:8086"

	// 设置查询参数
	values := url.Values{}
	values.Set("u", "admin")
	values.Set("p", "!Zyh123456")
	values.Set("q", "SHOW DATABASES")

	// 构造 GET 请求
	req, err := http.NewRequest("GET", influxdbURL+"/query?"+values.Encode(), nil)
	if err != nil {
		fmt.Println("创建请求失败：", err)
		return
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败：", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println("HTTP 状态码：", resp.StatusCode)
	fmt.Println("响应内容：")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(body))
}
