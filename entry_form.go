package main

import (
	"fmt"
	"strconv"

	"github.com/geistesk/procon/pc"

	"github.com/rivo/tview"
)

const pagesNameEntryForm = "entryform"

var (
	entryFormText  string
	entryFormType  string
	entryFormValue string
	entryFormEntry *pc.Entry
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
	entryFormText = text
	entryFormType = "Pro"
	if value < 0 {
		entryFormType = "Contra"
	}
	entryFormValue = formCreateValueText(value)
	entryFormEntry = nil

	var typeDropDownInit = 0
	if value < 0 {
		typeDropDownInit = 1
	}

	var form = tview.NewForm()

	form.AddInputField("Text", text, 0, nil,
		func(text string) {
			entryFormText = text
		})

	form.AddDropDown("Type", []string{"Pro", "Contra"}, typeDropDownInit,
		func(option string, _ int) {
			entryFormType = option
		})

	form.AddInputField("Value", formCreateValueText(value), 3,
		func(text string, _ rune) bool {
			n, err := strconv.Atoi(text)
			return err == nil && n > 0 && n <= 10
		}, func(text string) {
			entryFormValue = text
		})

	form.AddButton("Save", func() {
		// Already checked in value's accept function
		value, _ := strconv.Atoi(entryFormValue)
		if entryFormType == "Contra" {
			value *= -1
		}

		entry, err := pc.NewEntry(entryFormText, value)
		if err != nil {
			panic(err)
		}

		if entryFormEntry == nil {
			dataList.AddEntry(entry)
		} else {
			entryFormEntry.Text = entryFormText
			entryFormEntry.Value = value
		}

		redrawTable()

		pages.SwitchToPage(pagesNameTable)
		pages.RemovePage(pagesNameEntryForm)

		changed = true
		isTable = true
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage(pagesNameTable)
		pages.RemovePage(pagesNameEntryForm)

		isTable = true
	})

	form.SetBorder(true).SetTitle("New Entry").SetTitleAlign(tview.AlignLeft)

	return form
}

// newEntryFormFromEntry creates a new form for editing an entry.
func newEntryFormFromEntry(entry *pc.Entry) *tview.Form {
	form := newEntryForm(entry.Text, entry.Value)
	form.SetTitle(fmt.Sprintf("Edit %s", entry.Text))

	entryFormEntry = entry

	return form
}
