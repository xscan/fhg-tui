package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func Fetch(number int) []*ProjectItem {
	projects := make([]*ProjectItem, 0)

	if number < 1 {
		return projects
	}

	uri := fmt.Sprintf("https://hellogithub.com/periodical/volume/%d", number)

	c := colly.NewCollector()

	c.MaxDepth = 1

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Name)
		// re := regexp.MustCompile(`\d+`)
		// numberStr := re.FindString(e.Text)
		// number, _ = strconv.Atoi(numberStr)
		// 	fmt.Println(number)
	})

	c.OnHTML(".pb-4", func(e *colly.HTMLElement) {
		if e.Attr("id") != "" {

			category := e.DOM.Children().First().Text()
			// fmt.Println(category)
			e.ForEach(".pb-4", func(i int, h *colly.HTMLElement) {

				project := &ProjectItem{
					category: category,
					number:   strconv.Itoa(number),
				}

				header := h.DOM.Find(".mb-2.flex.flex-row.pt-3")
				// 获取标题 链接
				href := header.Find("a")
				val, _ := href.Attr("href")
				// fmt.Println(href.Text())
				// fmt.Println(val, err)

				project.uri = val
				project.name = href.Text()

				// 获取star fork
				starAndFork := header.Find(".text-sm")
				star := starAndFork.Children().First().Text()
				fork := starAndFork.Children().Last().Text()
				// fmt.Println(star, fork)

				// star = star.Rea
				project.star = strings.ReplaceAll(star, "Star ", "")
				project.fork = strings.ReplaceAll(fork, "Fork ", "")
				// 获取简介
				mdDom := h.DOM.Find(".markdown-body")
				desc := mdDom.Find("p").Text()
				// fmt.Println(desc)
				project.desc = desc

				projects = append(projects, project)
			})
		}

	})

	c.OnResponse(func(s *colly.Response) {
		// fmt.Println(string(s.Body))
	})

	c.Visit(uri)

	// for _, v := range projects {
	// 	fmt.Printf("%v\n", v)
	// }
	return projects
	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")

	// // Dump json to the standard output
	// enc.Encode(projects)
	// fmt.Println(uri)
}

// 获取站点信息
// 返回最大期数
// 返回项目数量
func FetchWebsite() map[string]interface{} {

	result := make(map[string]interface{}, 0)

	c := colly.NewCollector()

	c.MaxDepth = 1

	c.OnResponse(func(s *colly.Response) {
		// {"success":true,"repo_total":2613,"categories":[{"name":"C 项目"},{"name":"C# 项目"},{"name":"C++ 项目"},{"name":"CSS 项目"},{"name":"Go 项目"},{"name":"Java 项目"},{"name":"JavaScript 项目"},{"name":"Kotlin 项目"},{"name":"Objective-C 项目"},{"name":"PHP 项目"},{"name":"Python 项目"},{"name":"Ruby 项目"},{"name":"Rust 项目"},{"name":"Swift 项目"},{"name":"其它"},{"name":"开源书籍"},{"name":"机器学习"}],"volumes":[{"num":85},{"num":84},{"num":83},{"num":82},{"num":81},{"num":80},{"num":79},{"num":78},{"num":77},{"num":76},{"num":75},{"num":74},{"num":73},{"num":72},{"num":71},{"num":70},{"num":69},{"num":68},{"num":67},{"num":66},{"num":65},{"num":64},{"num":63},{"num":62},{"num":61},{"num":60},{"num":59},{"num":58},{"num":57},{"num":56},{"num":55},{"num":54},{"num":53},{"num":52},{"num":51},{"num":50},{"num":49},{"num":48},{"num":47},{"num":46},{"num":45},{"num":44},{"num":43},{"num":42},{"num":41},{"num":40},{"num":39},{"num":38},{"num":37},{"num":36},{"num":35},{"num":34},{"num":33},{"num":32},{"num":31},{"num":30},{"num":29},{"num":28},{"num":27},{"num":26},{"num":25},{"num":24},{"num":23},{"num":22},{"num":21},{"num":20},{"num":19},{"num":18},{"num":17},{"num":16},{"num":15},{"num":14},{"num":13},{"num":12},{"num":11},{"num":10},{"num":9},{"num":8},{"num":7},{"num":6},{"num":5},{"num":4},{"num":3},{"num":2},{"num":1}]}

		jsonString := string(s.Body)
		var data map[string]interface{}
		err := json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// fmt.Println(data["repo_total"].(float64))
		// fmt.Println(len(data["volumes"].([]interface{})))

		result["total"] = data["repo_total"].(float64)
		result["newVolumes"] = len(data["volumes"].([]interface{}))

	})
	uri := "https://api.hellogithub.com/v1/periodical/"
	c.Visit(uri)
	return result
}
