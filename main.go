package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		envFile   = flag.String("env-file", "", "Path to .env file")
		name      = flag.String("name", "app-secret", "Name of the Kubernetes Secret")
		namespace = flag.String("namespace", "default", "Namespace for the Kubernetes Secret")
	)
	flag.Parse()

	data := make(map[string]string)

	if *envFile != "" {
		err := loadEnvFile(*envFile, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
			os.Exit(1)
		}
	}

	loadEnvironmentVariables(data)

	if len(data) == 0 {
		fmt.Fprintf(os.Stderr, "No environment variables found\n")
		os.Exit(1)
	}

	generateSecretYAML(*name, *namespace, data)
}

func loadEnvFile(filename string, data map[string]string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		value = strings.Trim(value, `"'`)

		data[key] = value
	}

	return scanner.Err()
}

func loadEnvironmentVariables(data map[string]string) {
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			if _, exists := data[key]; !exists {
				data[key] = value
			}
		}
	}
}

func generateSecretYAML(name, namespace string, data map[string]string) {
	fmt.Printf("apiVersion: v1\n")
	fmt.Printf("kind: Secret\n")
	fmt.Printf("metadata:\n")
	fmt.Printf("  name: %s\n", name)
	fmt.Printf("  namespace: %s\n", namespace)
	fmt.Printf("type: Opaque\n")
	fmt.Printf("data:\n")

	for key, value := range data {
		encodedValue := base64.StdEncoding.EncodeToString([]byte(value))
		fmt.Printf("  %s: %s\n", key, encodedValue)
	}
}
