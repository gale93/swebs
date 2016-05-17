package filehandler

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//GetJSON will seek in json directory and serve the file requested
func GetJSON(path string) string {

	// get current path
	path, error := filepath.Abs(filepath.Dir(os.Args[0]))
	if error != nil {
		return ("[error]: " + error.Error())
	}

	dat, err := ioutil.ReadFile(path + "/resources/json/" + path + ".json")

	if err != nil {
		return "not_found: " + err.Error()
	}

	return string(dat)
}
