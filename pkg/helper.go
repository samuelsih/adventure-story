package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

//Adventure represents the JSON field for whole story adventure
type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

//Option represents the option data for story.json
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
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

// GetJSONData read the json file and return it
// return []byte because ReadAll() always return []byte
func GetStory(path string) (story map[string]Adventure) {
	file := OpenJsonFile(path)
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	return story
}