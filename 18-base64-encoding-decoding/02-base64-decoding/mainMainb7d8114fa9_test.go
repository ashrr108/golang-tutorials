package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func decodeBase64(encodedData string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}
	return string(decodedData), nil
}

func TestDecodeBase64(t *testing.T) {
	t.Run("Standard Base64 Decoding", func(t *testing.T) {
		encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
		expectedResult := "Gol@ng is Awesome?~"

		result, err := decodeBase64(encodedData)
		if err != nil {
			t.Errorf("Error decoding Base64 encoded data: %v", err)
		}

		if result != expectedResult {
			t.Errorf("Expected %s, got %s", expectedResult, result)
		}
	})

	t.Run("URL and filename-safe Base64 decoding", func(t *testing.T) {
		urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
		expectedResult := "Gol@ng is Awesome_~"

		result, err := decodeBase64(urlSafeBase64EncodedData)
		if err != nil {
			t.Errorf("Error decoding Base64 encoded data: %v", err)
		}

		if result != expectedResult {
			t.Errorf("Expected %s, got %s", expectedResult, result)
		}
	})

	t.Run("Invalid Base64 data", func(t *testing.T) {
		invalidEncodedData := "R29sQG5nIGlzIEF3ZXNvbWU?fg=="

		_, err := decodeBase64(invalidEncodedData)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}

func main() {
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
	decodedData, err := decodeBase64(encodedData)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(decodedData)

	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	urlSafeData, err := decodeBase64(urlSafeBase64EncodedData)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(urlSafeData)
}