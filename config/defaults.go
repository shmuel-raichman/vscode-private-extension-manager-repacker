// B"H
/*
Package config provides defaults config json and config structs
*/
package config

var defaultConfig []byte

// init Initialize default config json
func init() {

	defaultConfig = []byte(`{
		"ExtentionsIDs": [
			"ms-vscode.PowerShell", "redhat.vscode-yaml", "leizongmin.node-module-intellisense", "ms-vscode.vscode-typescript-next", "redhat.java"
		],
		"queryUrl": "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery"
	}`)
}

// GetDefaultConfig return default config as byte array
func GetDefaultConfig() []byte {
	return defaultConfig
}
