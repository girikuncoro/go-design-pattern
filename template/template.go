package main

import "fmt"

type Crawler interface {
	scrapeWebpage()
	parseHtml()
	store()
}

type AbstractCrawler struct {
	crawler Crawler
}

func (self *AbstractCrawler) Start() {
	self.crawler.scrapeWebpage()
	self.crawler.parseHtml()
	self.crawler.store()
}

type NewyorkTimesCrawler struct {
	*AbstractCrawler
}

func (self *NewyorkTimesCrawler) scrapeWebpage() {
	fmt.Println("Crawling newyork times webpage")
}

func (self *NewyorkTimesCrawler) parseHtml() {
	fmt.Println("Parse the scraped newyork times html page")
}

func (self *NewyorkTimesCrawler) store() {
	fmt.Println("Store processed newyork times page to DB")
}

type AljazeeraCrawler struct {
	*AbstractCrawler
}

func (self *AljazeeraCrawler) scrapeWebpage() {
	fmt.Println("Crawling aljazeera webpage")
}

func (self *AljazeeraCrawler) parseHtml() {
	fmt.Println("Parse the scraped aljazeera html page")
}

func (self *AljazeeraCrawler) store() {
	fmt.Println("Store processed aljazeera page to DB")
}

func main() {
	crawler := &AbstractCrawler{&NewyorkTimesCrawler{}}
	crawler.Start()

	crawler = &AbstractCrawler{&AljazeeraCrawler{}}
	crawler.Start()
}
