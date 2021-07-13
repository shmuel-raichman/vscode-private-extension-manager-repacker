package utils

import (
	"log"
	"os"
)

func CreateDir(dirname string) {
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDirWithAllContent(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Fatal(err)
	}
}
