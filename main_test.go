package main

import (
	"encoding/base64"
	"os"
	"strings"
	"testing"
)

func TestLoadEnvFile(t *testing.T) {
	content := `
# This is a comment
KEY1=value1
KEY2="value2"
KEY3='value3'
EMPTY_LINE=

# Another comment
KEY4=value4
`
	tmpFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	data := make(map[string]string)
	err = loadEnvFile(tmpFile.Name(), data)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"KEY1":       "value1",
		"KEY2":       "value2",
		"KEY3":       "value3",
		"EMPTY_LINE": "",
		"KEY4":       "value4",
	}

	if len(data) != len(expected) {
		t.Errorf("Expected %d keys, got %d", len(expected), len(data))
	}

	for key, expectedValue := range expected {
		if actualValue, exists := data[key]; !exists {
			t.Errorf("Expected key %s not found", key)
		} else if actualValue != expectedValue {
			t.Errorf("Expected value %s for key %s, got %s", expectedValue, key, actualValue)
		}
	}
}

func TestLoadEnvironmentVariables(t *testing.T) {
	os.Setenv("TEST_VAR1", "test_value1")
	os.Setenv("TEST_VAR2", "test_value2")
	defer func() {
		os.Unsetenv("TEST_VAR1")
		os.Unsetenv("TEST_VAR2")
	}()

	data := make(map[string]string)
	data["EXISTING_KEY"] = "existing_value"

	loadEnvironmentVariables(data)

	if value, exists := data["TEST_VAR1"]; !exists {
		t.Error("Expected TEST_VAR1 to be loaded")
	} else if value != "test_value1" {
		t.Errorf("Expected test_value1, got %s", value)
	}

	if value, exists := data["TEST_VAR2"]; !exists {
		t.Error("Expected TEST_VAR2 to be loaded")
	} else if value != "test_value2" {
		t.Errorf("Expected test_value2, got %s", value)
	}

	if value, exists := data["EXISTING_KEY"]; !exists {
		t.Error("Expected EXISTING_KEY to be preserved")
	} else if value != "existing_value" {
		t.Errorf("Expected existing_value, got %s", value)
	}
}

func TestBase64Encoding(t *testing.T) {
	testValue := "test-value"
	expected := base64.StdEncoding.EncodeToString([]byte(testValue))

	if !strings.Contains(expected, "dGVzdC12YWx1ZQ==") {
		t.Errorf("Base64 encoding not working as expected")
	}
}
