package logging

import (
	"encoding/json"
	"fmt"
	"go-toolkit/file"
	"go-toolkit/homedir"
	"log"
	"os"
)

const Path = "~/logs/logging.txt"

func PrintInfo() {
	conf := getConfiguration()
	fmt.Println(conf)
}

func getPath() string {
	expand, err := homedir.Expand(Path)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.MakeDirByFile(expand)
	if err != nil {
		log.Fatalln(err)
	}
	return expand
}

func getConfiguration() *configuration {
	open, err := os.Open("log_conf.json")
	if err != nil {
		path := getPath()
		return &configuration{Enabled: true, Path: path}
	}

	defer func(file *os.File) {
		if file != nil {
			err := file.Close()
			if err != nil {
				return
			}
		}
	}(open)

	conf := configuration{}
	decoder := json.NewDecoder(open)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Println(err)
	}
	return &conf
}

type configuration struct {
	Enabled bool
	Path    string
}
