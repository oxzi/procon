package pc

import (
	"errors"
	"fmt"
)

const valueAbsMax = 10

// Entry is one entry in a pros-cons list. The Text field describes its purpose
// and Value its characteristic on a scale from -10 to 10. Negative numbers are
// representing a "contra", positive numbers a "pro".
type Entry struct {
	_struct bool `codec:",toarray"`

	Text  string
	Value int
}

// NewEntry returns a new Entry. The value should be between -10 and 10.
func NewEntry(text string, value int) (e Entry, err error) {
	if value < -1*valueAbsMax || value > valueAbsMax {
		err = errors.New(fmt.Sprintf(
			"Value should be between -%d and %d", valueAbsMax, valueAbsMax))
		return
	}

	e = Entry{
		Text:  text,
		Value: value,
	}
	return
}

// IsPro determines if this Entry is a "pro" value.
func (e Entry) IsPro() bool {
	return e.Value > 0
}

// AbsValue returns the absolute Value.
func (e Entry) AbsValue() int {
	if e.Value < 0 {
		return -1 * e.Value
	} else {
		return e.Value
	}
}

// String returns the string representation of this Entry.
func (e Entry) String() string {
	var proConText string
	if e.IsPro() {
		proConText = "Pro"
	} else {
		proConText = "Con"
	}

	return fmt.Sprintf("%s: %s %d", e.Text, proConText, e.AbsValue())
}
