// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package structs

import "encoding/json"

func UnmarshalExtentionRequest(data []byte) (ExtentionRequest, error) {
	var r ExtentionRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtentionRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ExtentionRequest struct {
	Filters []Filter `json:"filters"`
	Flags   int64    `json:"flags"`
}

type Filter struct {
	Criteria []Criterion `json:"criteria"`
}

type Criterion struct {
	FilterType int64  `json:"filterType"`
	Value      string `json:"value"`
}
