package main

import "github.com/gdamore/tcell/v2"

var (
	Category2FetchCategory = map[string]string{
		"Java":            "Java 项目",
		"Python":          "Python 项目",
		"Javascript":      "Javascript 项目",
		"Rust":            "Rust 项目",
		"C":               "C 项目",
		"Cpp":             "C++ 项目",
		"Php":             "PHP 项目",
		"ObjectC":         "Objective-C 项目",
		"Go":              "Go 项目",
		"Css":             "Css 项目",
		"Csharp":          "C# 项目",
		"Kotlin":          "Kotlin 项目",
		"Swift":           "Swift 项目",
		"MachineLearning": "机器学习",
		"Ruby":            "Ruby 项目",
		"Book":            "开源书籍",
		"Other":           "其他",
	}
	FetchCategory2Category = map[string]string{
		"Java 项目":        "Java",
		"Python 项目":      "Python",
		"Javascript 项目":  "Javascript",
		"JavaScript 项目":  "Javascript",
		"Rust 项目":        "Rust",
		"C 项目":           "C",
		"C++ 项目":         "Cpp",
		"PHP 项目":         "Php",
		"Objective-C 项目": "ObjectC",
		"Go 项目":          "Go",
		"Css 项目":         "Css",
		"C# 项目":          "Csharp",
		"Kotlin 项目":      "Kotlin",
		"Swift 项目":       "Swift",
		"机器学习":           "MachineLearning",
		"Ruby 项目":        "Ruby",
		"开源书籍":           "Book",
		"其他":             "Other",
		"其它":             "Other",
	}

	TableCategoryColorMap = map[string]tcell.Color{
		"Java":            tcell.ColorRed,
		"Python":          tcell.Color100,
		"Javascript":      tcell.Color101,
		"Rust":            tcell.Color102,
		"C":               tcell.Color103,
		"Cpp":             tcell.Color104,
		"Php":             tcell.Color105,
		"ObjectC":         tcell.Color106,
		"Go":              tcell.Color107,
		"Css":             tcell.Color108,
		"Csharp":          tcell.Color109,
		"Kotlin":          tcell.Color110,
		"Swift":           tcell.Color111,
		"MachineLearning": tcell.Color112,
		"Ruby":            tcell.Color113,
		"Book":            tcell.Color114,
		"Other":           tcell.ColorTeal,
	}

	// 表格选择样式
	tableSelectedStyle = tcell.Style{}

	helpText = `CTRL j/k 切换 浏览/搜索 模式
搜索模式
Ctrl+h 获得帮助
输入 #{数字} 按期数搜索
输入 ${类别} 按类别搜索（未实现）
其他按关键字搜索（未实现）

浏览模式：
k/j 上/下 移动一行
u/d 上/下 移动五行
gg 移动至首行
G  移动至末行
h/l 前/后 翻页 （未实现）
o 查看（关闭）详细
s 帮 HG 点个小星星吧
ENTER 打开 GitHub 页面
q 退出应用`

// helpText = `CTRL j/k 切换 浏览/搜索 模式
// 搜索模式
// Ctrl+h 获得帮助
// 输入 #{数字} 按期数搜索
// 输入 ${类别} 按类别搜索
// 其他按关键字搜索

// 浏览模式：
// k/j 上/下 移动一行
// u/d 上/下 移动五行
// gg 移动至首行
// G  移动至末行
// h/l 前/后 翻页
// o 查看（关闭）详细
// s 帮 HG 点个小星星吧
// ENTER 打开 GitHub 页面
// q 退出应用`
)
