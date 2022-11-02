package main

import "github.com/meoww-bot/requests"

func main() {

	req := requests.Requests()

	resp, _ := req.Get("http://go.xiulian.net.cn", requests.Header{"Referer": "http://www.jeapedu.com"})

	println(resp.Text())

}
