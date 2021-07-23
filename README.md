## About
### Tool to repacke and publish vscode extensions for use with [vscode-private-extension-manager](https://github.com/joelspadin-garmin/vscode-private-extension-manager)

The motivation for this tools is because during my work for a few companies I encountered an issue where we use artifactory as mirror for standard's repositories such as npm, Maven, docker, etc. this way we can scan them before use inside the company network. 

Visual studio code extentions however don't have this option and is very useful tool for most of developers no netter what's language they write, but without extensions and with strict no download policy it make it harder to use it as you would when you have extensions and if you already downloaded extension and installed it you will probably never updated it because its too much Hassle.

So when I encounter joelspadin-garmin vscode-private-extention-manager I tried it and liked it, so I created this tool that that take a list of extensions IDs and download tham repack to npm and upload them to your own npm registry so you can use the privare-extention-manager with this repo, and keep it updated woth just single execution manualy or via simple cron job that keeps your extention mirror repo up to date. 

To create this tool I opend vscode developer-tools like in chrome browser and extracted the request it made wile i searched for extention with it ID.

What it's dose, for each extention ID request the extention manifest from it extarting the list of files urls requierd to repack the extention
- README.me
- icon.png
- CHANGELOG.md
- extension.vsix

Then downloads those files, and publish them to npm registry given in the config file. 

## Running
Make sure you logged in to your dedecated npm extention repo (NO scoped repo allowd)

Config file accespted pathes:
```bash
# Current execution dir:
repacker-config.json
# User home
# Linux:
USER_HOME/.vsix-to-npm/config.json
# Windows:
USER_HOME\.vsix-to-npm\config.json
```

Here example config
```json
{
	"ExtentionsIDs": [
		"ms-vscode.PowerShell", "redhat.vscode-yaml", "leizongmin.node-module-intellisense", "ms-vscode.vscode-typescript-next", "redhat.java"
	],
	"queryUrl": "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery",
	"registry": "http://localhost:4873/",
	"downloadPath": "ext"
}
```
Config is requierd since the default it complied with just the the above config.
```
./repacker
```

# Building
```bash
go build -o repacker vscode-ext
# Build on linux for Windows
env GOOS=windows GOARCH=amd64 go build -o repacker.exe vscode-ext
```


# Login to verdaccio
```
# Cretae user first time
npm adduser --registry http://repo:4873
# Login
npm login
```
