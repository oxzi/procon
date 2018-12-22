package proscons

import "testing"

func TestNewList(t *testing.T) {
	l := NewList("test", "test.pc")

	if l.Name != "test" {
		t.Errorf("List's name should be \"test\" instead of \"%s\"", l.Name)
	}
	if l.Filename != "test.pc" {
		t.Errorf("List's filename should be \"test.pc\" instead of \"%s\"", l.Filename)
	}
	if l.Entries != nil {
		t.Errorf("List's entries should be nil instead of \"%v\"", l.Entries)
	}
}

func TestAddEntry(t *testing.T) {
	l := NewList("test", "test.pc")

	if len(l.Entries) != 0 {
		t.Errorf("List's entires should be empty, instead of \"%v\"", l.Entries)
	}

	e, _ := NewEntry("Foo", 5)
	l.AddEntry(e)

	if len(l.Entries) != 1 {
		t.Errorf("List's entires should contain one element, instead of %d", len(l.Entries))
	}

	for i := 2; i <= 10; i++ {
		e, _ := NewEntry("Foo", 5)
		l.AddEntry(e)

		if len(l.Entries) != i {
			t.Errorf("List's entires should contain %d element, instead of %d", i, len(l.Entries))
		}
	}
}

func TestRemoveEntry(t *testing.T) {
	l := NewList("test", "test.pc")

	// Removing a non existing entry from an empty list should cause no problems.
	e, _ := NewEntry("Foo", 5)
	l.RemoveEntry(e)

	l.AddEntry(e)
	l.RemoveEntry(e)

	if len(l.Entries) != 0 {
		t.Errorf("List is not empty after deleting entry: %v", l.Entries)
	}

	l.AddEntry(e)
	l.RemoveEntry(Entry{"Foo", 5})

	if len(l.Entries) != 0 {
		t.Errorf("List is not empty after deleting entry: %v", l.Entries)
	}
}

func TestSumValues(t *testing.T) {
	l := NewList("test", "test.pc")

	if p, c := l.SumValues(); p != 0 || c != 0 {
		t.Errorf("SumValues is not 0, 0 for an empty list: %d, %d", p, c)
	}

	tests := []struct {
		entry Entry
		pros  int
		cons  int
	}{
		{Entry{"", 1}, 1, 0},
		{Entry{"", 1}, 2, 0},
		{Entry{"", -1}, 2, 1},
		{Entry{"", 10}, 12, 1},
		{Entry{"", -5}, 12, 6},
		{Entry{"", -3}, 12, 9},
	}

	for _, test := range tests {
		l.AddEntry(test.entry)

		if p, c := l.SumValues(); p != test.pros || c != test.cons {
			t.Errorf("SumValues returned %d, %d instead of %d, %d",
				p, c, test.pros, test.cons)
		}
	}
}

func TestProsConsEntries(t *testing.T) {
	l := NewList("test", "test.pc")

	if p, c := l.ProsConsEntries(); p != nil || c != nil {
		t.Errorf("ProsConsEntries is not nil, nil for an empty list: %v, %v", p, c)
	}

	entries := []Entry{
		Entry{"P1", 1},
		Entry{"P2", 5},
		Entry{"C1", -1},
		Entry{"C2", -5},
	}

	for _, entry := range entries {
		l.AddEntry(entry)
	}

	pros, cons := l.ProsConsEntries()

	if len(pros) != 2 || len(cons) != 2 {
		t.Errorf("Length of pros and cons list is wrong; %d, %d instead of 2, 2",
			len(pros), len(cons))
	}

	tests := []struct {
		eOriginal *Entry
		eGet      *Entry
	}{
		{&entries[1], pros[0]},
		{&entries[0], pros[1]},
		{&entries[3], cons[0]},
		{&entries[2], cons[1]},
	}

	for _, test := range tests {
		if *test.eOriginal != *test.eGet {
			t.Errorf("Values of ProsConsEntries are not equal to given ones: %v != %v",
				test.eOriginal, test.eGet)
		}
	}

	for i := 0; i < len(l.Entries); i++ {
		l.Entries[i].Text = "foo"
	}

	if pros[0].Text != "foo" {
		t.Errorf("Value has not changed: %v", pros[0])
	}

	for i := 0; i < len(pros)-1; i++ {
		if pros[i].AbsValue() < pros[i+1].AbsValue() {
			t.Errorf("Pros is not sorted")
		}
	}

	for i := 0; i < len(cons)-1; i++ {
		if cons[i].AbsValue() < cons[i+1].AbsValue() {
			t.Errorf("Cons is not sorted")
		}
	}
}
