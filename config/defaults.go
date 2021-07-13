// B"H
package config

var defaultConfig []byte

func init() {

	defaultConfig = []byte(`{
		"ExtentionsIDs": [
			"ms-vscode.PowerShell", "redhat.vscode-yaml", "leizongmin.node-module-intellisense", "ms-vscode.vscode-typescript-next", "redhat.java"
		],
		"queryUrl": "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery"
	}`)
}

func GetDefaultConfig() []byte {
	return defaultConfig
}
