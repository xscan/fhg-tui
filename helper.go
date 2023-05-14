package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func alert(pages *tview.Pages, id string, message string) *tview.Pages {
	model := tview.NewModal().
		SetText(message).
		AddButtons([]string{"确定"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.HidePage(id).RemovePage(id)
		}).SetBackgroundColor(tcell.ColorGray)

	return pages.AddPage(
		id,
		model,
		true,
		true,
	)
}

// MatchNumber 匹配期数字符串
func MatchNumber(input string) (match []string, err error) {
	re := regexp.MustCompile(`#(\d+)`)
	if re.Match([]byte(input)) {
		match = re.FindStringSubmatch(input)
		err = nil
	} else {
		err = errors.New("not match")
	}
	return
}

// openBrower 调用终端打开浏览器
func openBrower(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("open", url)
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func parseNumer2K(number float64) string {
	var result string
	c := number / 1000
	result = fmt.Sprintf("%0.1fK", c)
	return result
}
