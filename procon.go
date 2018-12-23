package main

import (
	"fmt"
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

// openFile opens or creates a List based on the first parameter.
func openFile() {
	if args := os.Args[1:]; len(args) == 1 {
		if err := loadDataList(args[0]); os.IsNotExist(err) {
			dataList = new(pc.List)
			dataList.Filename = args[0]
		} else if err != nil {
			fmt.Printf("procon, failed to open list: %v\n\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Usage: procon filename")
		fmt.Println("  filename will be opened or created\n")
		os.Exit(1)
	}
}

func main() {
	openFile()

	app = tview.NewApplication()
	pages = tview.NewPages()

	setupTable()
	pages.AddAndSwitchToPage(pagesNameTable, table, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			return nil
		}

		if isTable {
			tableHandleKeyPress(event)
		}

		return event
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("procon, failed to launch UI: %v\n\n", err)
		os.Exit(1)
	}
}
