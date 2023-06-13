package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := `{"alg":"HS256","typ":"JWT"}`

	// 使用标准编码进行 Base64 编码
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)

	de := "e1t9AuQ79PzVzXbRTUDNDjfOl+KKUIQjsOsOUj7q88A="
	// 解码Base64编码后的数据
	decoded, err := base64.StdEncoding.DecodeString(de)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
	//r := http.NewServeMux()
	//// 设置路由规则和处理函数
	//// ...
	//srv := &http.Server{
	//	Addr:    "10.121.12.61:9000",
	//	Handler: r,
	//}
	//
	//log.Printf("Starting server at %s...\n", srv.Addr)
	//if err := srv.ListenAndServe(); err != nil {
	//	log.Fatalf("Server failed: %v\n", err)
	//}
}
