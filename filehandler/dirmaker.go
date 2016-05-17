package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetUp will create the needed directories
func SetUp() error {

	errMsg := "Cannot create resources dir. Consider -checkdir=false.\n"

	// get current path
	path, error := filepath.Abs(filepath.Dir(os.Args[0]))
	if error != nil {
		return fmt.Errorf(errMsg + "[error] : " + error.Error())
	}

	if err := os.MkdirAll(path+"/resources", 0711); err != nil {
		return fmt.Errorf(errMsg + "[error] : " + err.Error())
	}

	if err := os.MkdirAll(path+"/resources/json", 0711); err != nil {
		return fmt.Errorf(errMsg + "[error] : " + err.Error())
	}

	f, err := os.Create(path + "/resources/index.html")
	if err != nil {
		return fmt.Errorf(errMsg + "[error] : " + err.Error())
	}

	f.Write([]byte("<h1>Hello simple world!</h1>"))

	return nil
}
