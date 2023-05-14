package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DetailView struct {
	item         *ProjectItem
	layoutDetail tview.Primitive
}

func NewDetailView(item *ProjectItem) *DetailView {
	view := &DetailView{
		item: item,
	}
	view.layoutDetail = view.layout()

	return view
}

func buildItem(title, value string, border bool) tview.Primitive {

	text := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetText(value).
		SetTextColor(tcell.ColorYellow)

	flex := tview.NewFlex().AddItem(text, 0, 1, false)
	flex.SetBorder(border).SetTitle(title).SetTitleAlign(tview.AlignLeft).SetTitleColor(tcell.ColorYellow).SetBorderColor(tcell.ColorYellow)

	return flex

}
func (d *DetailView) layout() tview.Primitive {

	first := tview.NewFlex().SetDirection(tview.FlexRowCSS).
		AddItem(buildItem(" 🐝 项目名称 ", d.item.name, true), 0, 1, false).
		AddItem(buildItem(" 🏁 项目地址 ", d.item.uri, true), 0, 1, false)

	star := tview.NewFlex().SetDirection(tview.FlexRowCSS).
		AddItem(buildItem("", fmt.Sprintf("🌟 Star: %s", d.item.star), true), 0, 1, false).
		AddItem(buildItem("", fmt.Sprintf("👀 Fork: %s", d.item.fork), true), 0, 1, false)

	content := buildItem(" 🍗 简介 ", d.item.desc, true)

	layout := tview.NewFlex().SetDirection(tview.FlexColumnCSS).
		AddItem(tview.NewBox(), 1, 1, false).
		AddItem(first, 3, 1, false).
		AddItem(star, 3, 1, false).
		AddItem(tview.NewBox(), 1, 1, false).
		AddItem(content, 0, 1, false).
		AddItem(tview.NewBox(), 1, 1, false)

	layout.SetBorder(true).SetTitle(" 项目详情 ")
	layout.SetTitleColor(tcell.ColorYellow).SetBorderColor(tcell.ColorYellow)
	layout.SetBorderPadding(0, 0, 1, 1)

	paddind := tview.NewFlex().SetDirection(tview.FlexColumnCSS).AddItem(layout, 0, 1, false)
	paddind.SetBorderPadding(6, 6, 4, 4)

	// .SetTitle("项目详情")
	// layout.SetBorderPadding(10, 10, 4, 4)

	// input := tview.NewTextView()
	// input.SetText("Hello World")

	// input1 := tview.NewTextView()
	// input1.SetText("Hello World11")

	// flex := tview.NewFlex().AddItem(input, 0, 1, false)
	// flex.SetBorder(true).SetTitle("标题")

	// flex1 := tview.NewFlex().AddItem(input1, 0, 1, false)
	// flex1.SetBorder(true).SetTitle("标题")

	// flexs := tview.NewFlex().SetDirection(tview.FlexRowCSS).AddItem(flex, 0, 1, false).AddItem(flex1, 0, 1, false)

	// flexmain := tview.NewFlex().SetDirection(tview.FlexColumnCSS).AddItem(flexs, 3, 1, false).AddItem(flexs, 3, 1, false)
	return paddind
}
