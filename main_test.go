package main

import (
	"testing"

	"github.com/florianherrengt/ksecli/models"
)

func mockReadValue(name string) string {
	return name
}

func TestHello(t *testing.T) {
	secrets := models.Secrets{}
	secrets.APIVersion = "v123"
	result := getSecretsValue(mockReadValue, secrets)
	if result.APIVersion != "v123" {
		t.Errorf("Unexpected APIVersion, expected: '%s', got:  '%s'", "v123", result.APIVersion)
	}
	if result.Data.DatabaseUsername != "Database username" {
		t.Errorf("Unexpected APIVersion, expected: '%s', got:  '%s'", "Database username", result.Data.DatabaseUsername)
	}
	if result.Data.DatabasePassword != "Database password" {
		t.Errorf("Unexpected APIVersion, expected: '%s', got:  '%s'", "Database password", result.Data.DatabasePassword)
	}
}
