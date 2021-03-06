package models

import (
	"encoding/json"
	"log"
	"os"
)

type Page struct {
	PageNo      int
	PageSize    int
	Next        int
	Previous    int
	TotalRecord int
	TotalPage   int
	Pages       []int
	URL         string
}

type Model struct {
	Number  int
	Name    string
	Job     string
	Click   int
	Page    string
	Address []string
}

type Config struct {
	DBuri string
}

func (config *Config) ReadConfig() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/src/config.json")
	if err != nil {
		log.Println("Open config file error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		log.Println("Decode config file error:", err)
	}
}

func (page *Page) CalculatePageInfo() {

	var previous, next, totalPage, start, end = page.PageNo - 1, page.PageNo + 1, 0, 0, 0

	if page.TotalRecord%page.PageSize == 0 {
		totalPage = page.TotalRecord / page.PageSize
	} else {
		totalPage = page.TotalRecord/page.PageSize + 1
	}

	if page.PageNo == 1 {
		previous = 1
	}
	if page.PageNo == page.TotalPage {
		next = page.TotalPage
	}

	if page.PageNo-2 <= 1 {
		start = 1
	} else {
		start = page.PageNo - 2
	}

	if page.PageNo+2 >= totalPage {
		end = totalPage
	} else {
		end = page.PageNo + 2
	}

	page.Previous = previous
	page.Next = next
	page.TotalPage = totalPage

	page.Pages = make([]int, end-start+1)
	for i := start; i <= end; i++ {
		page.Pages[i-start] = i
	}
}
