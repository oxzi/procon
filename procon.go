package main

import (
	"github.com/geistesk/procon/pc"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	var list = pc.NewList("Moving", "moving.pc")

	e1, _ := pc.NewEntry("Bigger flat", 6)
	e2, _ := pc.NewEntry("Better Job", 4)
	e3, _ := pc.NewEntry("No known people", -7)
	e4, _ := pc.NewEntry("Irksome", -5)

	list.AddEntry(e1)
	list.AddEntry(e2)
	list.AddEntry(e3)
	list.AddEntry(e4)

	var app = tview.NewApplication()
	var table = tview.NewTable().SetSeparator(tview.Borders.Vertical)

	table.SetCell(0, 0,
		tview.NewTableCell("Pros").
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 1,
		tview.NewTableCell("Cons").
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorYellow))

	pros, cons := list.ProsConsEntries()

	for i := 0; i < len(pros); i++ {
		table.SetCell(i+1, 0,
			tview.NewTableCell(pros[i].Text).SetAlign(tview.AlignRight))
	}

	for i := 0; i < len(cons); i++ {
		table.SetCell(i+1, 1,
			tview.NewTableCell(cons[i].Text).SetAlign(tview.AlignLeft))
	}

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
