package main

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

//Story make map of adventure
//Create this so we're not doing repetitive map[string]Adventure
type Story map[string]Adventure

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
func GetStory(path string) Story {
	var story Story

	file := OpenJsonFile(path)
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	return story
}

func main() {
	story := GetStory("story/story.json")

	fmt.Println(story)
}
