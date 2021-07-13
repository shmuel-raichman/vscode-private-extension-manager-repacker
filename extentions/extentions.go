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

func GetExtentionMeta(url string, method string, extensionId string) (*structs.ExtentionResaults, error) {
	// url := "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery"
	// method := "POST"

	// extId := "ms-vscode.PowerShell"

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

	payload, err := extentionRequest.Marshal()
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		return &structs.ExtentionResaults{}, err
	}
	req.Header.Add("accept", "application/json;api-version=3.0-preview.1")
	req.Header.Add("accept-language", "en-US")
	req.Header.Add("content-type", "application/json")
	// req.Header.Add("x-market-client-id", "VSCode 1.58.0-insider")

	res, err := client.Do(req)
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}

	extentionResaults, err := structs.UnmarshalExtentionResaults(body)
	if err != nil {
		return &structs.ExtentionResaults{}, err
	}
	return &extentionResaults, nil
}

func MakePackageJson(fileSource string, extensionId string, files [5]string) {
	/* First: declared map of string with empty interface
	which will hold the value of the parsed json. */
	var result map[string]interface{}
	/* Second: Unmarshal the json string string by converting
	it to byte into map */
	json.Unmarshal(utils.GetFileContent(fileSource), &result)
	/* Third: Read the Value by its key */
	result["private"] = false
	result["files"] = files
	result["icon"] = "media/icon.png"
	fmt.Println("Private :", result["private"])
	menifastJson, _ := json.MarshalIndent(result, "", " ")

	err := ioutil.WriteFile(fmt.Sprintf("%s/package.json", extensionId), menifastJson, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func DownloadExtentionFiles(baseExtentionsPath string, extensionID string, extentionResaults structs.ExtentionResaults) {

	var extentionPath string = filepath.Join(baseExtentionsPath, extensionID)

	utils.DeleteDirWithAllContent(extentionPath)

	utils.CreateDir(extentionPath)
	utils.CreateDir(filepath.Join(baseExtentionsPath, extensionID, "media"))

	fmt.Println("extentionPath: ", extentionPath)

	files := [5]string{"README.md", "media/icon.png", "package.josn", "extension.vsix", "CHANGELOG.md"}

	for _, file := range extentionResaults.Results[0].Extensions[0].Versions[0].Files {
		parts := strings.Split(file.AssetType, ".")
		switch parts[len(parts)-1] {
		case "Manifest":
			MakePackageJson(file.Source, extentionPath, files)
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
