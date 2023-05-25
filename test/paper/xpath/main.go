package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/tcxiaotudou/crawler/test/paper"
)

/**
XPath版本
*/

// var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	url := "https://www.thepaper.cn/"
	body, err := paper.Fetch(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}

	doc, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Println("htmlquery.Parse failed:%v", err)
	}

	nodes := htmlquery.Find(doc, `//div[@class="small_cardcontent__BTALp"]//h2`)

	for _, node := range nodes {
		fmt.Println("fetch card ", node.FirstChild.Data)
	}
}
