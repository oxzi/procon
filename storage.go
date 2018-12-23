package main

import (
	"errors"
	"os"

	"github.com/geistesk/procon/pc"
)

// saveDataList serializes the internal dataList to its filename.
func saveDataList() error {
	if dataList == nil {
		return errors.New("No data list exists")
	}

	f, err := os.Create(dataList.Filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = dataList.EncodeListToCbor(f); err != nil {
		return err
	}

	return nil
}

// loadDataList loads the given filename into the dataList.
func loadDataList(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	list, err := pc.DecodeListFromCbor(f)
	if err != nil {
		return err
	}

	list.Filename = filename
	dataList = &list
	return nil
}
