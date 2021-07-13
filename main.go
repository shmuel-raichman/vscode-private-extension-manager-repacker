package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"vscode-ext/structs"
	"vscode-ext/utils"
)

// type MyJsonName struct {
// 	Filters []struct {
// 		Criteria []struct {
// 			FilterType int64  `json:"filterType"`
// 			Value      string `json:"value"`
// 		} `json:"criteria"`
// 	} `json:"filters"`
// 	Flags int64 `json:"flags"`
// }

// type Welcome struct {
// 	Filters []Filter `json:"filters"`
// 	Flags   int64    `json:"flags"`
// }

// type Filter struct {
// 	Criteria []Criterion `json:"criteria"`
// }

// type Criterion struct {
// 	FilterType int64  `json:"filterType"`
// 	Value      string `json:"value"`
// }

func main() {

	url := "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery"
	method := "POST"

	extId := "ms-vscode.PowerShell"

	criterionId := structs.Criterion{
		FilterType: 7,
		Value:      extId,
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
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accept", "application/json;api-version=3.0-preview.1")
	req.Header.Add("accept-language", "en-US")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-market-client-id", "VSCode 1.58.0-insider")
	req.Header.Add("Cookie", "VstsSession=%7B%22PersistentSessionId%22%3A%22c21e54c9-d2b9-470f-905a-34cdcaf80780%22%2C%22PendingAuthenticationSessionId%22%3A%2200000000-0000-0000-0000-000000000000%22%2C%22CurrentAuthenticationSessionId%22%3A%2200000000-0000-0000-0000-000000000000%22%2C%22SignInState%22%3A%7B%7D%7D")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	extentionResaults, err := structs.UnmarshalExtentionResaults(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(extentionResaults.Results[0].Extensions[0].Publisher.PublisherName)
	// fmt.Println(extentionResaults.Results[0].Extensions[0].Publisher.DisplayName)
	// fmt.Println(extentionResaults.Results[0].Extensions[0].ExtensionName)
	// fmt.Println(extentionResaults.Results[0].Extensions[0].DisplayName)
	// fmt.Println(extentionResaults.Results[0].Extensions[0].Versions[0].Version)
	currentExtentionDir := "./ext-01"
	err = os.Mkdir(currentExtentionDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	
	files := [5]string{"README.md", "media/icon.png", "package.josn", "extension.vsix", "CHANGELOG.md"}

	for _, file := range extentionResaults.Results[0].Extensions[0].Versions[0].Files {
		parts := strings.Split(file.AssetType, ".")
		// println(parts[len(parts)-1])
		switch parts[len(parts)-1] {
		case "Manifest":

			/* First: declared map of string with empty interface
			which will hold the value of the parsed json. */
			var result map[string]interface{}
			/* Second: Unmarshal the json string string by converting
			it to byte into map */
			json.Unmarshal(getFile(file.Source), &result)
			/* Third: Read the Value by its key */
			result["private"] = false
			result["files"] = files
			result["icon"] = "media/icon.png"
			fmt.Println("Private :", result["private"])
			menifastJson, _ := json.MarshalIndent(result, "", " ")

			err = ioutil.WriteFile(fmt.Sprintf("%s/package.json", currentExtentionDir), menifastJson, 0644)
			if err != nil {
				fmt.Println(err)
			}
			println(file.Source)
		case "Details":
			utils.DownloadFile(fmt.Sprintf("%s/README.md", currentExtentionDir), file.Source, false)
			println(file.Source)
		case "Default":
			utils.DownloadFile(fmt.Sprintf("%s/media/icon.png", currentExtentionDir), file.Source, false)
			println(file.Source)
		case "Changelog":
			utils.DownloadFile(fmt.Sprintf("%s/CHANGELOG.md", currentExtentionDir), file.Source, false)
			println(file.Source)
		case "VSIXPackage":
			utils.DownloadFile(fmt.Sprintf("%s/extension.visx", currentExtentionDir), file.Source, false)
			println(file.Source)
		}
	}
}

func getFile(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
