package main

import "github.com/meoww-bot/requests"

func main() {

	data := requests.Datas{
		"name": "requests_post_test",
	}
	resp, _ := requests.Post("https://www.httpbin.org/post", data)
	println(resp.Text())
}
