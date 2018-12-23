package main

import (
	"github.com/rivo/tview"
)

const pagesNameQuitForm = "quitform"

// newQuitEntryForm is opened when the state is changed and asks the user to
// save or discard the changes.
func newQuitEntryForm() *tview.Form {
	form := tview.NewForm().
		AddButton("Save", func() {
			if err := saveDataList(); err != nil {
				panic(err)
			}

			app.Stop()
		}).
		AddButton("Discard", func() {
			app.Stop()
		})

	form.SetBorder(true).
		SetTitle("Discard changes?").
		SetTitleAlign(tview.AlignLeft)

	return form
}
