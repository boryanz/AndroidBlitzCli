package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	AppBuildGradle AppBuildGradle `json:"app_build_gradle"`
	Architecture   Architecture   `json:"architecture"`
}

type AppBuildGradle struct {
	PACKAGE_NAME        string `json:"package_name"`
	MINIMUM_SDK         int    `json:"minimum_sdk"`
	TARGET_SDK          int    `json:"target_sdk"`
	COMPILE_SDK         int    `json:"compile_sdk"`
	APP_VERSION         string `json:"app_version"`
	IS_MINIFIED_ENABLED bool   `json:"is_minified_enabled"`
}

type Architecture struct {
	IS_MVVM bool `json:"is_mvvm"`
}

func ParseConfigJson() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("Parsing of config.json failed %v: ", err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var config Config

	if err := json.Unmarshal(byteValue, &config); err != nil {
		fmt.Printf("Unmarshalling error: %v\n", err)
	}

	fmt.Println(config.AppBuildGradle.PACKAGE_NAME)

	return config
}
