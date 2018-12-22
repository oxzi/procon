package main

import (
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

func setupExampleList() {
	dataList = pc.NewList("Moving", "moving.pc")

	e1, _ := pc.NewEntry("Bigger flat", 6)
	e2, _ := pc.NewEntry("Whatever", 4)
	e3, _ := pc.NewEntry("Better Job", 4)
	e4, _ := pc.NewEntry("No known people", -10)
	e5, _ := pc.NewEntry("Irksome", -5)

	dataList.AddEntry(e1)
	dataList.AddEntry(e2)
	dataList.AddEntry(e3)
	dataList.AddEntry(e4)
	dataList.AddEntry(e5)
}

func main() {
	// XXX replace this with file handling
	setupExampleList()

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

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
