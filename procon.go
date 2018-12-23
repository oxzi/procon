package main

import (
	"os"

	"github.com/geistesk/procon/pc"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	dataList *pc.List

	app     *tview.Application
	pages   *tview.Pages
	isTable bool = true
)

func main() {
	if args := os.Args[1:]; len(args) == 1 {
		if err := loadDataList(args[0]); os.IsNotExist(err) {
			dataList = new(pc.List)
			dataList.Filename = args[0]
		} else if err != nil {
			panic(err)
		}
	} else {
		panic("One parameter hurr durr")
	}

	app = tview.NewApplication()
	pages = tview.NewPages()

	setupTable()
	pages.AddAndSwitchToPage(pagesNameTable, table, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if isTable {
			tableHandleKeyPress(event)
		}
		return event
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
