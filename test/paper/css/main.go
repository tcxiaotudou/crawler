package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tcxiaotudou/crawler/test/paper"
)

/**
CSS版本
*/

// var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {

	url := "https://www.thepaper.cn/"
	body, err := paper.Fetch(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}

	doc.Find("div.small_cardcontent__BTALp h2").Each(func(i int, s *goquery.Selection) {
		// 获取匹配标签中的文本
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
