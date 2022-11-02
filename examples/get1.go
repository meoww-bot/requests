package main

import "github.com/meoww-bot/requests"

func main() {

	resp, _ := requests.Get("http://go.xiulian.net.cn")
	println(resp.Text())
}
