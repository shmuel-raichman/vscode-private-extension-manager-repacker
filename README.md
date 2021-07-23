## About
### Tool to repacke and publish vscode extensions for use with [vscode-private-extension-manager](https://github.com/joelspadin-garmin/vscode-private-extension-manager)

The motivation for this tools is because during my work few companies i encountered an issue where we use artifactory as mirror for standard's repositories such as npm, Maven, docker, etc. this way we can scan them before use inside the company network. 

Visual studio code is very useful tool most of the developers no netter what's language they write, but without extensions and with strict no download policy it make it harder to use it as you would when you have extensions and if you already downloaded extension and installed it you will probably never updated it because its too much Hussle.

So when i encounter joelspadin-garmin vscode-private-extention-manager i tried it and liked it very much, so i created this tool that that take a lost of extensions IDs and download tham repack to npm and upload them to your own npm registry so you can use the privare-extention-manager with this repo, and keep it updated woth just single execution. 

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


Login to verdaccio
```
npm adduser --registry http://repo:4873
```
