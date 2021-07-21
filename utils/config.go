// B"H
/*
Package utils needs more comments
*/
package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	"errors"
	"path/filepath"
	"vscode-ext/config"

	// Config
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

var (
	configData config.Config
)

func init() {
	readConfigs()
}

func ReadConfigJson(jsonConfigBytesBuffer *bytes.Buffer) config.Config { //structs.Config {

	var configObject config.Config

	decoder := json.NewDecoder(jsonConfigBytesBuffer)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&configObject); err != nil {
		panic(err)
	}

	return configObject
}

func useDefaultConfig() *bytes.Buffer {

	defaultBinariesBytesBufferConfig := bytes.NewBuffer(config.GetDefaultConfig())
	return defaultBinariesBytesBufferConfig

}

func readConfigFile(configFilePath string) *bytes.Buffer {
	log.Println("Reading ", configFilePath, "config")

	jsonConfigFile, err := os.Open(configFilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonConfigFile.Close()

	// read our opened File as a byte array.
	jsonFileByteValue, _ := ioutil.ReadAll(jsonConfigFile)
	jsonConfigBytesBuffer := bytes.NewBuffer(jsonFileByteValue)

	return jsonConfigBytesBuffer
}

// func readConfigs() { // config.Config{
// 	homeDirPath, err := os.UserHomeDir()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defaultConfigPath := filepath.Join(homeDirPath, ".vsix-to-npm", "config.json")
// 	var configFilePath string = defaultConfigPath
// 	// var config config.Config
// 	_, err = exec.LookPath(configFilePath)
// 	// If no error then file exist else download file.
// 	if err == nil {
// 		log.Println("Read config from file.")
// 		configData = ReadConfigJson(readConfigFile(configFilePath))
// 	} else {
// 		log.Println("Use default config.")
// 		configData = ReadConfigJson(useDefaultConfig())
// 	}
// 	// return config
// }

func readConfigs() {

	var configFilePath string
	var configFromFile bool

	// Get user home dir
	homeDirPath, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		homeDirPath = ""
	}

	// Get current working diratory.
	currentWorkingDiractory, err := os.Getwd()
	if err != nil {
		log.Println(err)
		currentWorkingDiractory = ""
	}

	defaultConfigPath := filepath.Join(currentWorkingDiractory, "repacker-config.json")
	homeConfigPath := filepath.Join(homeDirPath, ".vsix-to-npm", "config.json")

	log.Println("Looking for config file in: ", homeConfigPath)
	_, err = os.Stat(homeConfigPath)
	if !errors.Is(err, fs.ErrNotExist) {
		configFilePath = homeConfigPath
		configFromFile = true
	}

	log.Println("Looking for config file in: ", defaultConfigPath)
	_, err = os.Stat(defaultConfigPath)
	if !errors.Is(err, fs.ErrNotExist) {
		configFilePath = defaultConfigPath
		configFromFile = true
	}

	if configFromFile {
		log.Println("Read config from: ", configFilePath)
		configData = ReadConfigJson(readConfigFile(configFilePath))
	} else {
		log.Println("")
		log.Println("Config file not found, using default config")
		configData = ReadConfigJson(useDefaultConfig())
	}
}

func GetExtentionsIDsConfig() []string {
	return configData.ExtentionsIDs
}

func GetQueryUrlConfig() string {
	return configData.QueryUrl
}

func GetRegistryConfig() string {
	return configData.RegistryUrl
}

func GetConfig() config.Config {
	return configData
}
