package filehandler

import "io/ioutil"

//GetJSON will seek in json directory and serve the file requested
func GetJSON(path string) string {
	dat, err := ioutil.ReadFile("resources/json/" + path + ".json")

	if err != nil {
		return "not_found: " + err.Error()
	}

	return string(dat)
}
