/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func (receiver *JsonConfig) Load() *JsonConfig {
	if CheckConfigFile() {
		JsonFile, ErrReadJsonFile := ioutil.ReadFile(CheckOsForConfig())
		if ErrReadJsonFile != nil {
			log.Fatalf("JSON JSON struct ErrReadJsonFile: " + ErrReadJsonFile.Error())
		}

		ErrUnmarshalJsonFile := json.Unmarshal(JsonFile, receiver)
		if ErrUnmarshalJsonFile != nil {
			log.Fatalf("JSON JSON struct ErrUnmarshalJsonFile: " + ErrUnmarshalJsonFile.Error())
		}

		return receiver
	}

	return nil
}

func CheckConfigFile() bool {
	file, err := os.OpenFile(CheckOsForConfig(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {

		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Can't load configuration | l.55 - config/config.go")
		}
	}(file)

	File, _ := file.Stat()
	if File.Size() == 0 {
		fmt.Println("Please, fill config file. Located at -> config.json")

		os.Exit(1)

		return false
	}

	return true
}

// UNUSED FUNCTION
func CheckConfigDirectory() bool {
	path := "config"

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)

		if err != nil {
			log.Println(err)
			os.Exit(1)

			return false
		}
	}

	return true
}

func CheckOsForConfig() string {
	var FilePath = ""

	if runtime.GOOS == "windows" {
		FilePath = FilePathWin
	} else {
		FilePath = FilePathUnix
	}

	return FilePath
}
