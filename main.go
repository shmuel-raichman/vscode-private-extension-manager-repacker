// B"H
/*
Package main
*/
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

	// For each extention id in config file
	// Request manifest into struct and downlaod each extention files
	for _, extentionID := range extensionsIDs {
		fmt.Println("extentionID: ", extentionID)
		// Get current extention manifest to ExtentionResaults struct
		extentionResaults, err := extentions.GetExtentionMeta(url, method, extentionID)
		if err != nil {
			log.Fatal(err)
		}

		// Download requierd extention files.
		extentions.DownloadExtentionFiles(baseExtentionsPath, extentionID, *extentionResaults)

		currentExtentionDir := filepath.Join(baseExtentionsPath, extentionID)
		// npm publish command parameters
		command := []string{
			"publish",
			currentExtentionDir,
			"--registry",
			utils.GetRegistryConfig(),
		}

		// Run npm publish command
		utils.ExecuteCommand("npm", command)
	}
}
