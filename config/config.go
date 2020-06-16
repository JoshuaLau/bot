package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"bufio"
	"strings"
)

var (
	Token string
	BotPrefix string
	Companies map[string]struct{} = make(map[string]struct{})

	config *configStruct
)

type configStruct struct {
	Token string  `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	fmt.Println("Reading from companies text file...")

	companiesFile, err := os.Open("./companies.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer companiesFile.Close()

	scanner := bufio.NewScanner(companiesFile)
	for scanner.Scan() {
		Companies[strings.ToLower(scanner.Text())] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

