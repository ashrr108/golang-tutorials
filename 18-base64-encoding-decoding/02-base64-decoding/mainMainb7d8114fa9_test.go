package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func decodeBase64(encodedData string, urlSafe bool) (string, error) {
	var decodedData []byte
	var err error

	if urlSafe {
		decodedData, err = base64.URLEncoding.DecodeString(encodedData)
	} else {
		decodedData, err = base64.StdEncoding.DecodeString(encodedData)
	}

	if err != nil {
		return "", err
	}
	return string(decodedData), nil
}

func TestDecodeBase64(t *testing.T) {
	// Test case 1: Standard Base64 decoding
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
	expectedResult := "Gol@ng is Awesome?~"
	result, err := decodeBase64(encodedData, false)
	if err != nil {
		t.Errorf("Error decoding Base64 encoded data %v", err)
	}
	if result != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, result)
	}

	// Test case 2: URL and filename-safe Base64 decoding
	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	expectedUrlSafeResult := "Gol@ng is Awesome_~"
	urlSafeResult, err := decodeBase64(urlSafeBase64EncodedData, true)
	if err != nil {
		t.Errorf("Error decoding Base64 encoded data %v", err)
	}
	if urlSafeResult != expectedUrlSafeResult {
		t.Errorf("Expected %s, got %s", expectedUrlSafeResult, urlSafeResult)
	}
}

func TestDecodeBase64InvalidInput(t *testing.T) {
	// Test case 3: Invalid input for Standard Base64 decoding
	invalidEncodedData := "R29sQG5nIGlzIEF3ZXNvbWU?@"
	_, err := decodeBase64(invalidEncodedData, false)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Test case 4: Invalid input for URL and filename-safe Base64 decoding
	invalidUrlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU?@"
	_, err = decodeBase64(invalidUrlSafeBase64EncodedData, true)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func main() {
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
	decodedData, err := decodeBase64(encodedData, false)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(decodedData)

	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	urlSafeData, err := decodeBase64(urlSafeBase64EncodedData, true)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(urlSafeData)
}