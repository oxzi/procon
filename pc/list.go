package pc

import (
	"reflect"
	"sort"
)

// List is a pros-cons list with a name, a filename and its entries.
type List struct {
	Name     string
	Filename string `codec:"-"`
	Entries  []Entry
}

// NewList creates a new List with the given parameters.
func NewList(name, filename string) *List {
	return &List{
		Name:     name,
		Filename: filename,
	}
}

// forEach applies the given function to each entry.
func (l *List) forEach(f func(*Entry, int)) {
	for i := 0; i < len(l.Entries); i++ {
		f(&l.Entries[i], i)
	}
}

// AddEntry adds an entry to the List's list of entries.
func (l *List) AddEntry(entry Entry) {
	l.Entries = append(l.Entries, entry)
}

// RemoveEntry removes the entry form the list of entries.
func (l *List) RemoveEntry(entry Entry) {
	var pos = -1
	l.forEach(func(e *Entry, i int) {
		if reflect.DeepEqual(*e, entry) {
			pos = i
		}
	})

	if pos != -1 {
		l.Entries = append(l.Entries[:pos], l.Entries[pos+1:]...)
	}
}

// SumValues returns the sum of all pros' and cons' values.
func (l *List) SumValues() (pros, cons int) {
	l.forEach(func(e *Entry, _ int) {
		if e.IsPro() {
			pros += e.AbsValue()
		} else {
			cons += e.AbsValue()
		}
	})
	return
}

// ProsConsEntries returns one slice of pros and one of cons entry pointers.
// Each slice is sorted by each entry's absolute value.
func (l *List) ProsConsEntries() (pros, cons []*Entry) {
	l.forEach(func(e *Entry, _ int) {
		if e.IsPro() {
			pros = append(pros, e)
		} else {
			cons = append(cons, e)
		}
	})

	var less = func(i, j int, list []*Entry) bool {
		return list[i].AbsValue() > list[j].AbsValue()
	}

	sort.Slice(pros, func(i, j int) bool { return less(i, j, pros) })
	sort.Slice(cons, func(i, j int) bool { return less(i, j, cons) })

	return
}
