// B"H
/*
Package main
*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vscode-ext/extentions"
	"vscode-ext/utils"
)

func main() {

	url := utils.GetConfig().QueryUrl
	method := "POST"
	var baseExtentionsPath string = utils.GetConfig().DownloadPath

	extensionsIDs := utils.GetConfig().ExtentionsIDs

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

		// Get current working diratory.
		currentWorkingDiractory, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		currentExtentionDir := filepath.Join(currentWorkingDiractory, baseExtentionsPath, extentionID)
		// npm publish command parameters
		npmPublishCommand := []string{
			"publish",
			// currentExtentionDir,
			"--registry",
			utils.GetConfig().RegistryUrl,
		}

		npmUnpublishCommand := []string{
			"unpublish",
			// currentExtentionDir,
			"--registry",
			utils.GetConfig().RegistryUrl,
			"--force",
		}

		println(currentExtentionDir)
		utils.ExecuteCommand("npm", npmUnpublishCommand, currentExtentionDir)
		utils.ExecuteCommand("npm", npmPublishCommand, currentExtentionDir)
	}
}
