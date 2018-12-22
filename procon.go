package main

import (
	"bytes"
	"fmt"

	"github.com/geistesk/procon/pc"
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

	var buff = new(bytes.Buffer)
	list.EncodeListToCbor(buff)

	fmt.Printf("%X\n", buff)

	var l2, _ = pc.DecodeListFromCbor(buff)
	fmt.Println(l2)
}
