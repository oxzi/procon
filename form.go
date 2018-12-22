package main

import (
	"fmt"
	"strconv"

	"github.com/geistesk/procon/pc"

	"github.com/rivo/tview"
)

const pagesNameForm = "form"

var (
	formText  string
	formType  string
	formValue string
)

// formCreateValueText creates a value string from an int.
func formCreateValueText(value int) (valueText string) {
	if value == 0 {
		valueText = ""
	} else if value < 0 {
		valueText = fmt.Sprintf("%d", -1*value)
	} else {
		valueText = fmt.Sprintf("%d", value)
	}
	return
}

// newEntryForm returns a new Form to create an entry.
func newEntryForm(text string, value int) *tview.Form {
	var typeDropDownInit = 0
	if value < 0 {
		typeDropDownInit = 1
	}

	var form = tview.NewForm()

	form.AddInputField("Text", text, 0, nil,
		func(text string) {
			formText = text
		})

	form.AddDropDown("Type", []string{"Pro", "Contra"}, typeDropDownInit,
		func(option string, _ int) {
			formType = option
		})

	form.AddInputField("Value", formCreateValueText(value), 3,
		func(text string, _ rune) bool {
			n, err := strconv.Atoi(text)
			return err == nil && n > 0 && n <= 10
		}, func(text string) {
			formValue = text
		})

	form.AddButton("Save", func() {
		// Already checked in value's accept function
		value, _ := strconv.Atoi(formValue)
		if formType == "Contra" {
			value *= -1
		}

		entry, err := pc.NewEntry(formText, value)
		if err != nil {
			panic(err)
		}

		dataList.AddEntry(entry)

		redrawTable()

		pages.SwitchToPage(pagesNameTable)
		pages.RemovePage(pagesNameForm)

		isTable = true
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage(pagesNameTable)
		pages.RemovePage(pagesNameForm)

		isTable = true
	})

	form.SetBorder(true).SetTitle("New Entry").SetTitleAlign(tview.AlignLeft)

	return form
}
