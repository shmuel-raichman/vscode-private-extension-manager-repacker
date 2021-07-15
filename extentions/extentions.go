// B"H
/*
Package extensions provides functions to handle vscode extentions requests, downloads and making extention Manifest a valid package.json
*/
package extentions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"vscode-ext/structs"
	"vscode-ext/utils"
)

// GetExtentionMeta calls vscode marketplace with extention id and return extention Manifest as pre configured struct
// Each faild step will return empty "*structs.ExtentionResaults" and err
func GetExtentionMeta(url string, method string, extensionId string) (*structs.ExtentionResaults, error) {
	// url := "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery"
	// method := "POST"

	// extId := "ms-vscode.PowerShell"

	/*
		Construct request objects to be marshel into json request body
		Hard coded values from sniffing vscode request
		The objects here are nested final object type "ExtentionRequest"
	*/
	// The only modified value here is "extensionId"
	criterionId := structs.Criterion{
		FilterType: 7,
		Value:      extensionId,
	}
	criterionInternal := structs.Criterion{
		FilterType: 8,
		Value:      "Microsoft.VisualStudio.Code",
	}

	filter := structs.Filter{
		Criteria: []structs.Criterion{criterionInternal, criterionId},
	}

	extentionRequest := structs.ExtentionRequest{
		Filters: []structs.Filter{filter},
		Flags:   950,
	}

	// Marshel the request object into json (marshel is object method)
	payload, err := extentionRequest.Marshal()
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}

	// Create http request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}

	// Headers from sniffing vscode request (maybe there is more headers to ommit)
	req.Header.Add("accept", "application/json;api-version=3.0-preview.1")
	req.Header.Add("accept-language", "en-US")
	req.Header.Add("content-type", "application/json")
	// req.Header.Add("x-market-client-id", "VSCode 1.58.0-insider")

	// Do the http request
	res, err := client.Do(req)
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}
	// https://tour.golang.org/flowcontrol/12
	// Not sure need further reading what this line
	// In general and whats is defer in particular
	defer res.Body.Close()

	// Convert results.body to byte array
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}

	// Unmarshel json into per defined "ExtentionResaults"
	extentionResaults, err := structs.UnmarshalExtentionResaults(body)
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}
	// If all this is cool let's return extention results
	// Basicly what I need from this struct is
	// Results[0].Extensions[0].Versions[0].Files which contains Urls for the extention files manifest, readme.md etc.
	// There is probably a better more error safe way to get it but this works for now.
	return &extentionResaults, nil
}

// MakePackageJson gets extention manifest url and extentionId(for output file path)
// Retrives the manifest and makes the following changes
// * Add - files for vscode private extention manager (README.md, extention.vsix, etc. )
// * Change - private field to false
// * Add - icon path
func MakePackageJson(fileSource string, extensionId string) {
	files := []string{"README.md", "media/icon.png", "package.josn", "extension.vsix", "CHANGELOG.md"}
	/* First: declared map of string with empty interface
	which will hold the value of the parsed json. */
	var result map[string]interface{}
	/* Second: Unmarshal the json string string by converting
	it to byte into map */
	json.Unmarshal(utils.GetFileContent(fileSource), &result)
	/* Third: Read the Value by its key */
	// Chnage or add values
	result["private"] = false
	result["files"] = files
	result["icon"] = "media/icon.png"

	// Not sure what going on here but it's working ;)
	menifastJson, _ := json.MarshalIndent(result, "", " ")

	// Recreate the manifest as packege.json file
	err := ioutil.WriteFile(fmt.Sprintf("%s/package.json", extensionId), menifastJson, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// DownloadExtentionFiles downloads extention meta files for repacking
func DownloadExtentionFiles(baseExtentionsPath string, extensionID string, extentionResaults structs.ExtentionResaults) {

	// use path.filepath to make it run on linux and windows
	var extentionPath string = filepath.Join(baseExtentionsPath, extensionID)

	// Delete extenions dir with all content, to avoid path exist issue and previus versions
	utils.DeleteDirWithAllContent(extentionPath)
	// Recreate extenions dir
	utils.CreateDir(extentionPath)
	// Recreate extenions icon dir
	utils.CreateDir(filepath.Join(baseExtentionsPath, extensionID, "media"))
	fmt.Println("extentionPath: ", extentionPath)

	// Range extention files and download them to extentionPath
	for _, file := range extentionResaults.Results[0].Extensions[0].Versions[0].Files {
		// Here is the file AssetType:
		// "Microsoft.VisualStudio.Services.VSIXPackage"
		// spliting it by dot "." and use switch on last part of AssetType
		parts := strings.Split(file.AssetType, ".")
		switch parts[len(parts)-1] {
		case "Manifest":
			MakePackageJson(file.Source, extentionPath)
			// println(file.Source)
		case "Details":
			utils.DownloadFile(fmt.Sprintf("%s/README.md", extentionPath), file.Source, false)
			// println(file.Source)
		case "Default":
			utils.DownloadFile(fmt.Sprintf("%s/media/icon.png", extentionPath), file.Source, false)
			// println(file.Source)
		case "Changelog":
			utils.DownloadFile(fmt.Sprintf("%s/CHANGELOG.md", extentionPath), file.Source, false)
			// println(file.Source)
		case "VSIXPackage":
			utils.DownloadFile(fmt.Sprintf("%s/extension.vsix", extentionPath), file.Source, false)
			// println(file.Source)
		}
	}
}
