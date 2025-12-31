package main

import (
	"encoding/base64"
	"fmt"
)

func bm(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
func jm(s string) (string, error) {
	date, error := base64.StdEncoding.DecodeString(s)
	return string(date), error
}
func main() {
	var choice int
	var ru string
	for {
		fmt.Println("1.编码/2.解码/3.结束")
		fmt.Scan(&choice)
		if choice == 3 {
			fmt.Println("结束！")
			break
		}
		switch choice {
		case 1:
			fmt.Println("请输入内容：")
			fmt.Scan(&ru)
			result := bm(ru)
			fmt.Println(result)
		case 2:
			fmt.Println("请输入内容：")
			fmt.Scan(&ru)
			result, err := jm(ru)
			if err != nil {
				fmt.Println("error")
			}
			fmt.Println(result)
		default:
			fmt.Println("无效选项")
		}
	}
}
