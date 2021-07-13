// B"H
package main

import (
	"fmt"
	"log"
	"path/filepath"
	"vscode-ext/extentions"
	"vscode-ext/utils"
)

func main() {

	url := utils.GetQueryUrlConfig()
	method := "POST"
	var baseExtentionsPath string = "ext"

	extensionsIDs := utils.GetExtentionsIDsConfig()

	for _, ext := range extensionsIDs {
		fmt.Println(ext)
		extentionResaults, err := extentions.GetExtentionMeta(url, method, ext)
		if err != nil {
			log.Fatal(err)
		}

		extentions.DownloadExtentionFiles(baseExtentionsPath, ext, *extentionResaults)
		currentExtentionDir := filepath.Join(baseExtentionsPath, ext)
		command := []string{
			"publish",
			currentExtentionDir,
			"--registry",
			"http://localhost:4873/",
		}

		utils.ExecuteCommand("npm", command)
	}
}
