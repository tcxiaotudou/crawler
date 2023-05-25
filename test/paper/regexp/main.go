package main

import (
	"fmt"
	"github.com/tcxiaotudou/crawler/test/paper"
	"regexp"
)

/**
正则表达式版本
*/

// var headerRe = regexp.MustCompile(`<div class="news_li"[\s\S]*?<h2>[\s\S]*?<a.*?target="_blank">([\s\S]*?)</a>`)
var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	url := "https://www.thepaper.cn/"
	body, err := paper.Fetch(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	matches := headerRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		fmt.Println("fetch card news:", string(m[1]))
	}
}
