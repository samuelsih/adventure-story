package main

import (
	"fmt"
	"os"
)


// const port = ":8080"

func OpenJsonFile(path string) *os.File {
	//open json file
	jsonFile, err := os.Open(path)

	//if error, print the error
	if err != nil {
		fmt.Println(err)
	}

	//defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return jsonFile
}

func main() {
	jsonFile := OpenJsonFile("tools/story.json")

	fmt.Println("Success reading jsonFile => " + jsonFile.Name())

}
