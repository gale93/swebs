package filehandler

import (
	"fmt"
	"os"
)

// SetUp will create the needed directories
func SetUp() error {

	errMsg := "Impossible create resources dir. Consider -checkdir=false.\n"

	if err := os.MkdirAll("resources", 0711); err != nil {
		return fmt.Errorf(errMsg + "[error] : " + err.Error())
	}

	if err := os.MkdirAll("resources/json", 0711); err != nil {
		return fmt.Errorf(errMsg + "[error] : " + err.Error())
	}

	f, err := os.Create("resources/index.html")
	if err != nil {
		return fmt.Errorf(errMsg + "[error] : " + err.Error())
	}

	f.Write([]byte("<h1>Hello simple world!</h1>"))

	return nil
}
