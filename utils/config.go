// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	"path/filepath"
	"vscode-ext/config"

	// Config
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func readConfigs() { // config.Config{
	homeDirPath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	defaultConfigPath := filepath.Join(homeDirPath, ".vsix-to-npm", "config.json")
	var configFilePath string = defaultConfigPath
	// var config config.Config
	_, err = exec.LookPath(configFilePath)
	// If no error then file exist else download file.
	if err == nil {
		log.Println("Read config from file.")
		configData = ReadConfigJson(readConfigFile(configFilePath))
	} else {
		log.Println("Use default config.")
		configData = ReadConfigJson(useDefaultConfig())
	}
	// return config
}

func GetExtentionsIDsConfig() []string {
	return configData.ExtentionsIDs
}

func GetQueryUrlConfig() string {
	return configData.QueryUrl
}
