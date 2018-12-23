package main

import (
	"fmt"

	"github.com/geistesk/procon/pc"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	columnPros int = 0
	columnCons int = 1

	pagesNameTable       string = "table"
	pagesNameDeleteModal string = "deletemodal"
	pagesNameQuitModal   string = "quitmodal"
)

var (
	table            *tview.Table
	tblPros, tblCons []*pc.Entry
)

// newDeleteEntryModal asks the user if the given entry should be deleted.
func newDeleteEntryModal(entry *pc.Entry) *tview.Modal {
	return tview.NewModal().
		SetText(fmt.Sprintf("Do you want to delete\n\"%s\"?", entry.Text)).
		AddButtons([]string{"Delete", "Cancel"}).
		SetDoneFunc(func(_ int, buttonLabel string) {
			if buttonLabel == "Delete" {
				dataList.RemoveEntry(*entry)
				redrawTable()

				changed = true
			}

			pages.SwitchToPage(pagesNameTable)
			pages.RemovePage(pagesNameDeleteModal)

			isTable = true
		})
}

// newQuitEntryModal is opened when the state is changed and asks the user to
// save or discard the changes.
func newQuitEntryModal() *tview.Modal {
	return tview.NewModal().
		SetText("You have unsaved changes.").
		AddButtons([]string{"Save", "Discard"}).
		SetDoneFunc(func(_ int, buttonLabel string) {
			if buttonLabel == "Save" {
				if err := saveDataList(); err != nil {
					panic(err)
				}
			}

			app.Stop()
		})
}

// entryRepresentation returns a representing string for the table.
func entryRepresentation(entry *pc.Entry) string {
	if entry.IsPro() {
		return fmt.Sprintf("[white]%s [green]%2d ", entry.Text, entry.AbsValue())
	} else {
		return fmt.Sprintf("[red] %-2d [white]%s", entry.AbsValue(), entry.Text)
	}
}

// setupTableHeader creats the table's header.
func setupTableHeader() {
	cols := []struct {
		no   int
		text string
	}{
		{columnPros, "Pros"},
		{columnCons, "Cons"},
	}

	for _, col := range cols {
		table.SetCell(0, col.no,
			tview.NewTableCell(col.text).
				SetSelectable(false).
				SetAlign(tview.AlignCenter).
				SetTextColor(tcell.ColorYellow))
	}
}

// syncListToTable draws all elements of the dataList to the table.
func syncListToTable() {
	if dataList == nil {
		return
	}

	table.SetTitle(dataList.Filename)

	tblPros, tblCons = dataList.ProsConsEntries()

	for i := 0; i < len(tblPros); i++ {
		table.SetCell(i+1, columnPros,
			tview.NewTableCell(entryRepresentation(tblPros[i])).
				SetAlign(tview.AlignRight))
	}

	for i := 0; i < len(tblCons); i++ {
		table.SetCell(i+1, columnCons,
			tview.NewTableCell(entryRepresentation(tblCons[i])).
				SetAlign(tview.AlignLeft))
	}
}

// redrawTable redraws the table.
func redrawTable() {
	table.Clear()

	setupTableHeader()
	syncListToTable()
}

// selectedTableEntry returns an optional pointer to the currently selected
// entry or nil otherwise.
func selectedTableEntry() *pc.Entry {
	var entry *pc.Entry

	r, c := table.GetSelection()
	if c == columnPros {
		if r-1 >= len(tblPros) {
			return nil
		}
		entry = tblPros[r-1]
	} else {
		if r-1 >= len(tblCons) {
			return nil
		}
		entry = tblCons[r-1]
	}

	return entry
}

// tableHandleKeyPress is called when the table is in focus and a key was pressed.
func tableHandleKeyPress(event *tcell.EventKey) {
	if event.Key() != tcell.KeyRune {
		return
	}

	switch event.Rune() {
	case 'x':
		if entry := selectedTableEntry(); entry != nil {
			isTable = false
			pages.AddAndSwitchToPage(
				pagesNameDeleteModal, newDeleteEntryModal(entry), true)
		}

	case 'a':
		_, pos := table.GetSelection()
		if pos == columnCons {
			pos = -1
		}

		isTable = false
		pages.AddAndSwitchToPage(pagesNameEntryForm, newEntryForm("", pos), true)

	case 'w':
		if err := saveDataList(); err != nil {
			panic(err)
		}
		changed = false

	case 'q':
		if changed {
			isTable = false
			pages.AddAndSwitchToPage(pagesNameQuitModal, newQuitEntryModal(), true)
		} else {
			app.Stop()
		}
	}
}

// setupTable creates the pros and cons table.
func setupTable() {
	table = tview.NewTable().
		SetSeparator(tview.Borders.Vertical).
		SetFixed(1, 1).
		SetSelectable(true, true).
		SetSelectedFunc(func(_, _ int) {
			if entry := selectedTableEntry(); entry != nil {
				isTable = false
				pages.AddAndSwitchToPage(
					pagesNameEntryForm, newEntryFormFromEntry(entry), true)
			}
		})

	table.SetBorder(true).SetTitleAlign(tview.AlignLeft)

	setupTableHeader()
	syncListToTable()
}
