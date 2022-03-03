package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadNotes() []map[string]string {

	// Building path to the notes.json file
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "notes.json")

	// Opening file
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return make([]map[string]string, 0)
	}

	// Closing the file in the end
	defer jsonFile.Close()

	// Getting slice of notes from file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]string
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func SaveNotes(notes []map[string]string) error {

	// Building path to the notes.json file
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "notes.json")

	// Getting jsonString from slice
	jsonString, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	// Writing to json file
	err2 := ioutil.WriteFile(path, jsonString, os.ModePerm)
	if err2 != nil {
		return err2
	}

	return nil
}
