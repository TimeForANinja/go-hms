package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func readFile() []byte {
	absPath, err := filepath.Abs("./util/querys.sql")
	if err != nil {
		panic(fmt.Errorf("Error getting filepath: \"%s\"", err))
	}
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(fmt.Errorf("Error loading querys: \"%s\"", err))
	}
	return data
}
