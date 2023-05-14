package main

import (
	"fmt"
	"os"
	"testing"
)

func TestKey(t *testing.T) {
	a := 'q'
	fmt.Println(a)
}

func TestMatchSearchNumber(t *testing.T) {
	input := "#2323ss"
	m, err := MatchNumber(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

func TestDB(t *testing.T) {
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	file := "/cache.db"
	filepath := fmt.Sprintf("%s%s", dir, file)
	db := DB{}
	init := db.Init(filepath)
	if init != nil {
		fmt.Println("初始化错误", init)
		return
	}
	item := ProjectItem{
		name:     "test1",
		number:   "2",
		category: "C",
		desc:     "test desc",
		uri:      "http://www.baidu.com",
		star:     "2201",
		fork:     "253",
	}
	db.Save(item)

	result := db.All()
	fmt.Println(result)
}

func Test2K(t *testing.T) {
	number := 1234.23
	c := number / 1000
	fmt.Printf("%0.1fK", c)
}
func TestFetchWebsite(t *testing.T) {
	info := FetchWebsite()
	fmt.Println(info)
}
func TestFetchData(t *testing.T) {
	a := Fetch(85)
	fmt.Println(len(a))
}
func TestFetch(t *testing.T) {

	// // Dump json to the standard output
	// enc.Encode(projects)
	// fmt.Println(uri)
}
