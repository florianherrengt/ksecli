package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/florianherrengt/ksecli/models"

	"gopkg.in/yaml.v2"
)

func main() {
	secrets := models.Secrets{}

	data, err := ioutil.ReadFile("secrets.yml")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	err = yaml.Unmarshal([]byte(data), &secrets)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if secrets.Data.DatabaseUsername == "" {
		fmt.Print("Database username: ")
		var input string
		fmt.Scanln(&input)
		secrets.Data.DatabaseUsername = b64.StdEncoding.EncodeToString([]byte(input))
	}

	if secrets.Data.DatabasePassword == "" {
		fmt.Print("Database password: ")
		var input string
		fmt.Scanln(&input)
		secrets.Data.DatabasePassword = b64.StdEncoding.EncodeToString([]byte(input))
	}

	d, err := yaml.Marshal(&secrets)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("secrets.yml", d, 0644)
}
