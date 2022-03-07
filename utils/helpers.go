package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Note struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

func LoadNotes() []*Note {

	// Building path to the notes.json file
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "notes.json")

	// Opening file
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return make([]*Note, 0)
	}

	// Closing the file in the end
	defer jsonFile.Close()

	// Getting slice of notes from file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []*Note
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func SaveNotes(notes []*Note) error {

	// Building path to the notes.json file
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "notes.json")

	// Getting jsonString from slice
	jsonString, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	// Writing to json file
	err = ioutil.WriteFile(path, jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
