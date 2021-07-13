// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package structs

import "encoding/json"

func UnmarshalExtentionResaults(data []byte) (ExtentionResaults, error) {
	var r ExtentionResaults
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtentionResaults) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ExtentionResaults struct {
	Results []Result `json:"results"`
}

type Result struct {
	Extensions     []Extension       `json:"extensions"`
	PagingToken    interface{}       `json:"pagingToken"`
	ResultMetadata []ResultMetadatum `json:"resultMetadata"`
}

type Extension struct {
	Publisher        Publisher   `json:"publisher"`
	ExtensionID      string      `json:"extensionId"`
	ExtensionName    string      `json:"extensionName"`
	DisplayName      string      `json:"displayName"`
	Flags            string      `json:"flags"`
	LastUpdated      string      `json:"lastUpdated"`
	PublishedDate    string      `json:"publishedDate"`
	ReleaseDate      string      `json:"releaseDate"`
	ShortDescription string      `json:"shortDescription"`
	Versions         []Version   `json:"versions"`
	Categories       []string    `json:"categories"`
	Tags             []string    `json:"tags"`
	Statistics       []Statistic `json:"statistics"`
	DeploymentType   int64       `json:"deploymentType"`
}

type Publisher struct {
	PublisherID   string `json:"publisherId"`
	PublisherName string `json:"publisherName"`
	DisplayName   string `json:"displayName"`
	Flags         string `json:"flags"`
}

type Statistic struct {
	StatisticName string  `json:"statisticName"`
	Value         float64 `json:"value"`
}

type Version struct {
	Version          string     `json:"version"`
	Flags            string     `json:"flags"`
	LastUpdated      string     `json:"lastUpdated"`
	Files            []File     `json:"files"`
	Properties       []Property `json:"properties"`
	AssetURI         string     `json:"assetUri"`
	FallbackAssetURI string     `json:"fallbackAssetUri"`
}

type File struct {
	AssetType string `json:"assetType"`
	Source    string `json:"source"`
}

type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResultMetadatum struct {
	MetadataType  string         `json:"metadataType"`
	MetadataItems []MetadataItem `json:"metadataItems"`
}

type MetadataItem struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
