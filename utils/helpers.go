package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadNotes() []map[string]string {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "notes.json")
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return make([]map[string]string, 0)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

func SaveNotes(notes []map[string]string) error {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "notes.json")
	jsonString, err := json.Marshal(notes)
	if err != nil {
		return err
	}
	err2 := ioutil.WriteFile(path, jsonString, os.ModePerm)
	if err2 != nil {
		return err2
	}
	return nil
}
