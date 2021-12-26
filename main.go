package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Story represents the JSON field
type Story struct {
	Title   string            `json:"title"`
	Story   []string          `json:"story"`
	Options map[string]string `json:"options"`
}


//OpenJsonFile open the json file
//return *File
func OpenJsonFile(path string) *os.File {
	//open json file
	jsonFile, err := os.Open(path)

	//if error, print the error
	if err != nil {
		fmt.Println(err)
	}

	return jsonFile
}

func main() {
	//get json file
	jsonFile := OpenJsonFile("tools/story.json")

	//ioutil.ReadAll() read all the jsonFile data
	//return it as a []byte
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	var story map[string]Story

	//Unmarshal parses the JSON-encoded data and stores it in story variable
	json.Unmarshal(byteValue, &story)

	for key, value := range story {
		fmt.Println(key)
		fmt.Println("title = " + value.Title)

		for _, s := range value.Story {
			fmt.Println(s)
		}
	}
}
