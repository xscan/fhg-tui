package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type mode int8

const (
	SearchMode mode = iota //搜索模式
	ListMode               //列表模式
	DetailMode             //详情模式
	HelpMode               //搜索模式
)

type AppView struct {
	app         *tview.Application
	title       *tview.TextView
	subTitle    *tview.TextView
	searchInput *tview.InputField
	pages       *tview.Pages
	helpModel   *tview.Modal
	list        *tview.List
	table       *tview.Table
	status      *tview.Flex

	mode           *tview.TextView
	currentMode    mode
	cacheDataModel map[int][]*ProjectItem
	contents       []*ProjectItem
	currentNumber  int
	maxNumber      int
}

func NewAppView(app *tview.Application, currentNumber int, maxNunber int) *AppView {
	view := &AppView{
		app:            app,
		list:           tview.NewList(),
		pages:          tview.NewPages(),
		title:          tview.NewTextView(),
		subTitle:       tview.NewTextView(),
		searchInput:    tview.NewInputField(),
		helpModel:      tview.NewModal(),
		table:          tview.NewTable(),
		mode:           tview.NewTextView(),
		contents:       make([]*ProjectItem, 0),
		currentMode:    SearchMode,
		cacheDataModel: make(map[int][]*ProjectItem),
		currentNumber:  currentNumber,
		maxNumber:      maxNunber,
	}

	view.title.SetText("HelloGithub").SetTextColor(tcell.ColorYellow).SetTextAlign(tview.AlignCenter)

	view.subTitle.SetText("分享 GitHub 上有趣、入门级的开源项目").SetTextColor(tcell.ColorYellow).SetTextAlign(tview.AlignCenter)

	// 搜索栏输入框
	view.searchInput.SetBorder(true)
	inputStyle := tcell.Style{}.Foreground(tcell.ColorWhite)
	view.searchInput.SetFieldStyle(inputStyle)
	view.searchInput.SetDoneFunc(func(key tcell.Key) {
		// 这是跳转匹配
		// 跳转期数
		input := view.searchInput.GetText()
		m, err := MatchNumber(input)
		if err == nil {
			// 跳转到期数
			number, _ := strconv.Atoi(m[1])
			view.gotoPage(number)
			// 设置模式
			view.app.SetFocus(view.table)
			view.currentMode = ListMode
			return
		} else {
			// 本地搜索，不线上搜
		}

		view.searchInput.SetText("")
	})

	view.contents = Fetch(view.currentNumber)
	// 添加到缓存
	view.cacheDataModel[view.currentNumber] = view.contents

	// 渲染表格
	view.renderTable()

	// view.title.SetText(strconv.Itoa(len(contents)))
	// 添加表格内容

	view.table.Select(0, 0)
	view.table.SetSelectable(true, false).SetSelectedFunc(func(row, column int) {
		// 选择行事件
	})

	// 设置选中颜色
	view.table.SetSelectedStyle(tableSelectedStyle)
	// 固定列表头
	view.table.SetFixed(1, 1)
	// view.table.SetEvaluateAllRows(true)

	// 设置表格边框
	view.table.SetBorder(true).SetTitle(" [yellow]搜索结果 ")

	view.table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	view.helpModel.SetBorder(true).SetTitle(" 帮助 ")

	view.helpModel.SetText("j/k 上下移动")

	m := FetchWebsite()

	//状态栏
	s := tview.NewFlex().
		SetDirection(tview.FlexRowCSS).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText("⏰"), 3, 1, false).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText("[yellow]2023-05-12 10:12:57"), 20, 1, false).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText("🌟"), 2, 1, false).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText(fmt.Sprintf("[yellow]%s", parseNumer2K(m["total"].(float64)))), 6, 1, false).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText("📚"), 2, 1, false).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText(fmt.Sprintf("[yellow]项目数 %d 个", int(m["total"].(float64)))), 22, 1, false)

	view.status = tview.NewFlex().SetDirection(tview.FlexRowCSS).
		AddItem(tview.NewTextView().SetDynamicColors(true).
			SetRegions(true).
			SetWordWrap(true).
			SetChangedFunc(func() {
				app.Draw()
			}).SetText("[white]按 [green]ctrl h [white]查看帮助 按 [green]q [white]键退出").SetTextAlign(tview.AlignLeft), 0, 1, false).
		AddItem(tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText(" [yellow]搜索模式 ").SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(s, 0, 1, false)

	searchInputLayout := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(view.searchInput, 0, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)
	// searchBox := tview.NewFrame(searchInputLayout)

	// 创建整体布局
	flex := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(view.title, 1, 1, false).
		AddItem(view.subTitle, 1, 1, false).
		AddItem(tview.NewBox(), 1, 1, false).
		AddItem(searchInputLayout, 3, 1, false).
		AddItem(view.table, 0, 1, false).
		AddItem(view.status, 1, 1, false)

	flex.SetBorderPadding(1, 2, 2, 2)

	view.pages.AddPage("main", flex, true, true)

	return view
}

// 上一页
func (a *AppView) prevPage() {
	a.currentNumber--
	a.gotoPage(a.currentNumber)
}

func (a *AppView) gotoPage(number int) {
	a.currentNumber = number
	if v, ok := a.cacheDataModel[a.currentNumber]; ok {
		a.contents = v
	} else {
		a.contents = Fetch(a.currentNumber)
		a.cacheDataModel[a.currentNumber] = a.contents
	}

	a.renderTable()
}

// 下一页
func (a *AppView) nextPage() {
	a.currentNumber++
	a.gotoPage(a.currentNumber)
}

// 渲染表格
func (a *AppView) renderTable() {

	a.table.Clear()

	// 表格初始化
	columnHeader := []string{
		"№", "名称", "期数", "分类", "介绍",
	}

	for i := 0; i < len(columnHeader); i++ {
		columnName := columnHeader[i]
		var cell *tview.TableCell

		if columnName == "介绍" {
			cell = tview.NewTableCell(fmt.Sprintf("[yellow]%s", columnName)).
				SetTextColor(tview.Styles.TitleColor).
				SetMaxWidth(0).SetExpansion(2)
		} else {
			cell = tview.NewTableCell(fmt.Sprintf("[yellow]%s", columnName)).
				SetTextColor(tview.Styles.TitleColor).
				SetMaxWidth(30).SetExpansion(1)
		}
		a.table.SetCell(0, i, cell)
	}

	var textColor tcell.Color
	var displayCategory string
	var ok bool
	for i := 1; i <= len(a.contents); i++ {
		item := a.contents[i-1]
		row := i
		// 如果存在映射分类
		if displayCategory, ok = FetchCategory2Category[item.category]; !ok {
			displayCategory = item.category
		}

		// 颜色
		if v, ok := TableCategoryColorMap[displayCategory]; ok {
			textColor = v
		} else {
			textColor = TableCategoryColorMap["Other"]
		}
		a.table.SetCell(row, 0, tview.NewTableCell(strconv.Itoa(i)).SetTextColor(textColor))
		a.table.SetCell(row, 1, tview.NewTableCell(item.name).SetTextColor(textColor))
		a.table.SetCell(row, 2, tview.NewTableCell(item.number).SetTextColor(textColor))
		a.table.SetCell(row, 3, tview.NewTableCell(displayCategory).SetTextColor(textColor))
		a.table.SetCell(row, 4, tview.NewTableCell(item.desc).SetTextColor(textColor))
	}
}

func main() {
	// 获取最新期数
	m := FetchWebsite()

	newNumber := int(m["newVolumes"].(int))

	app := tview.NewApplication()

	view := NewAppView(app, newNumber, newNumber)

	timer := time.NewTicker(time.Duration(time.Second * 1))

	go func(timer *time.Ticker) {
		for {
			<-timer.C
			// 更新时间
			tiemKey := view.status.GetItem(2).(*tview.Flex).GetItem(1).(*tview.TextView)
			tiemKey.SetDynamicColors(true).SetRegions(true).SetWordWrap(true).SetText(fmt.Sprintf("[yellow]%s", time.Now().Format("2006-01-02 15:04:05")))
			app.Draw()
		}
	}(timer)

	// 创建应用程序
	app.SetRoot(view.pages, true)
	// 设置全局事件
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlK {
			app.SetFocus(view.searchInput)
			a := view.status.GetItem(1).(*tview.TextView)
			a.SetDynamicColors(true).SetText(" [yellow]搜索模式 ")

			view.currentMode = SearchMode
			return nil
		}
		if event.Key() == tcell.KeyCtrlJ {
			app.SetFocus(view.table)
			view.currentMode = ListMode
			a := view.status.GetItem(1).(*tview.TextView)
			a.SetDynamicColors(true).SetText(" [yellow]列表模式 ")

			return nil
		}
		if event.Key() == tcell.KeyCtrlH {
			alert(view.pages, "alert", helpText)
			view.currentMode = HelpMode
			return nil
		}
		if event.Rune() == 'q' {
			app.Stop()
		}

		// 搜索模式
		if view.currentMode == SearchMode {
			view.subTitle.SetText(fmt.Sprintf("%v", event.Rune()))
			return event
		}

		// 列表模式模式
		if view.currentMode == ListMode {
			if event.Rune() == 'h' {
				// 上一页
				// app.Stop()
				if view.currentNumber > 1 {
					view.prevPage()
					// view.subTitle.SetText("上一页")
					a := view.status.GetItem(1).(*tview.TextView)
					a.SetDynamicColors(true).SetText(fmt.Sprintf("[yellow] ⇦ h   第 %d 期   l ⇨ ", view.currentNumber))
				}

				return nil
			}
			if event.Rune() == 'l' {
				// 下一页
				// app.Stop()
				if view.currentNumber < view.maxNumber {
					view.nextPage()
					// view.subTitle.SetText("下一页")
					a := view.status.GetItem(1).(*tview.TextView)
					a.SetDynamicColors(true).SetText(fmt.Sprintf("[yellow] ⇦ h   第 %d 期   l ⇨ ", view.currentNumber))
				}
				return nil
			}
			if event.Rune() == 'o' {
				// 打开详情
				row, _ := view.table.GetSelection()
				if row <= 0 {
					return nil
				}

				item := view.contents[row-1]
				detail := NewDetailView(item)

				view.pages.AddAndSwitchToPage(item.name, detail.layoutDetail, true)
				view.currentMode = DetailMode
				// view.subTitle.SetText("打开详情")
				return nil
			}
			if event.Key() == tcell.KeyEnter {
				// 打开链接
				row, _ := view.table.GetSelection()
				if row <= 0 {
					return nil
				}
				uri := view.contents[row-1].uri
				openBrower(uri)
				return nil
			}
			if event.Rune() == 'u' {
				// 表格上移动5行
				row, _ := view.table.GetSelection()
				newRowIndex := row - 5
				if newRowIndex <= 0 {
					newRowIndex = 0
				}

				view.table.Select(newRowIndex, 0)
				return nil
			}
			if event.Rune() == 'd' {
				// 表格下移5行
				row, _ := view.table.GetSelection()

				newRowIndex := row + 5
				if newRowIndex >= view.table.GetRowCount() {
					newRowIndex = view.table.GetRowCount() - 1
				}

				view.table.Select(newRowIndex, 0)
				return nil
			}

			return event
		}

		// 帮助模式
		if view.currentMode == HelpMode {
			// view.app.SetFocus(view.table)
			// view.currentMode = ListMode
			return event
		}

		if view.currentMode == DetailMode {
			row, _ := view.table.GetSelection()
			item := view.contents[row-1]
			view.pages.HidePage(item.name).RemovePage(item.name)
			view.pages.SwitchToPage("main")
			view.app.SetFocus(view.table)
			view.currentMode = ListMode
			return event
		}
		return event
	})
	// 启用鼠标
	// app.EnableMouse(true)

	// 设置焦点到搜索框
	app.SetFocus(view.searchInput)
	// 运行应用程序
	if err := app.Run(); err != nil {
		panic(err)
	}
}
