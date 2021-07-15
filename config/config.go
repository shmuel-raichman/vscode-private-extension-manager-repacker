// B"H
/*
Package config provides defaults config json and config structs
*/
package config

// Config contains variables that might need to be supplied
type Config struct {
	ExtentionsIDs []string `json:"ExtentionsIDs"`
	QueryUrl      string   `json:"queryUrl"`
	RegistryUrl   string   `json:"registry"`
}
