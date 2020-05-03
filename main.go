package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/florianherrengt/ksecli/models"

	"gopkg.in/yaml.v2"
)

// ValueReader gets the value for a secret
type ValueReader func(description string) string

func readValue(description string) string {
	fmt.Printf("%s: ", description)
	var input string
	fmt.Scanln(&input)
	return b64.StdEncoding.EncodeToString([]byte(input))
}

func getSecretsValue(readValue ValueReader, secrets models.Secrets) models.Secrets {
	if secrets.Data.DatabaseUsername == "" {
		secrets.Data.DatabaseUsername = readValue("Database username")
	}

	if secrets.Data.DatabasePassword == "" {
		secrets.Data.DatabasePassword = readValue("Database password")
	}

	return secrets
}

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

	secrets = getSecretsValue(readValue, secrets)
	d, err := yaml.Marshal(&secrets)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("secrets.yml", d, 0644)
}
