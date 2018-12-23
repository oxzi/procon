package main

import (
	"fmt"

	"github.com/geistesk/procon/pc"

	"github.com/rivo/tview"
)

const pagesNameDeleteForm = "deleteform"

// newDeleteEntryForm asks the user if the given entry should be deleted.
func newDeleteEntryForm(entry *pc.Entry) *tview.Form {
	form := tview.NewForm().
		AddButton("Delete", func() {
			dataList.RemoveEntry(*entry)
			redrawTable()

			pages.SwitchToPage(pagesNameTable)
			pages.RemovePage(pagesNameDeleteForm)

			changed = true
			isTable = true
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage(pagesNameTable)
			pages.RemovePage(pagesNameDeleteForm)

			isTable = true
		})

	form.SetBorder(true).
		SetTitle(fmt.Sprintf("Delete %s", entry.Text)).
		SetTitleAlign(tview.AlignLeft)

	return form
}
