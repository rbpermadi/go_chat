package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	Port        int
	HostName    string
	LogFilePath string
}

func loadConfig() Configuration {
	fmt.Println(os.Getenv("PORT"))
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println(err)
	}

	configuration := Configuration{
		Port:        port,
		HostName:    os.Getenv("HOST_NAME"),
		LogFilePath: os.Getenv("LOGFILE_PATH"),
	}

	fmt.Println(configuration)
	return configuration
}

func setupLogging(configuration Configuration) {
	nf, err := os.Create(configuration.LogFilePath)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func LoadConfigAndSetUpLogging() Configuration {
	configuration := loadConfig()
	setupLogging(configuration)

	return configuration
}
