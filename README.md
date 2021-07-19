## About
### Tool to repacke and publish vscode extensions for use with [vscode-private-extension-manager](https://github.com/joelspadin-garmin/vscode-private-extension-manager)

This tool is made by getting vscode network request when filtering extensions in the extensions tab.
> This is still an beta	 version 

## Running
Make sure you logged in to your dedecated npm extention repo (NO scoped repo allowd)

Put json config file in: <br>
`USER_HOME/.vsix-to-npm/config.json` <br> 

Or windows <br>
`USER_HOME\.vsix-to-npm\config.json`

Here example config
```json
{
	"ExtentionsIDs": [
		"ms-vscode.PowerShell",
		"redhat.vscode-yaml", 
		"leizongmin.node-module-intellisense", 
		"ms-vscode.vscode-typescript-next", 
		"redhat.java",
		"ms-python.python"
	],
	"queryUrl": "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery",
	"registry": "http://localhost:4873/"
}
```
Config is requierd since the default it complied with just the the above extentions IDs
```
./repacker
```

# Building
```bash
go build -o repacker vscode-ext
# Build on linux for Windows
env GOOS=windows GOARCH=amd64 go build -o repacker.exe vscode-ext
```