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

func decodeURLSafeBase64(encodedData string) (string, error) {
	decodedData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}
	return string(decodedData), nil
}

func TestDecodeBase64(t *testing.T) {
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
	expectedDecodedData := "Gol@ng is Awesome?"

	decodedData, err := decodeBase64(encodedData)
	if err != nil {
		t.Errorf("Error decoding Base64 encoded data: %v", err)
	}

	if decodedData != expectedDecodedData {
		t.Errorf("Expected %q, got %q", expectedDecodedData, decodedData)
	}
}

func TestDecodeBase64InvalidInput(t *testing.T) {
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg==="

	_, err := decodeBase64(encodedData)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDecodeURLSafeBase64(t *testing.T) {
	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	expectedDecodedData := "Gol@ng is Awesome?"

	decodedData, err := decodeURLSafeBase64(urlSafeBase64EncodedData)
	if err != nil {
		t.Errorf("Error decoding Base64 encoded data: %v", err)
	}

	if decodedData != expectedDecodedData {
		t.Errorf("Expected %q, got %q", expectedDecodedData, decodedData)
	}
}

func TestDecodeURLSafeBase64InvalidInput(t *testing.T) {
	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg==="

	_, err := decodeURLSafeBase64(urlSafeBase64EncodedData)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func main() {
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="

	decodedData, err := decodeBase64(encodedData)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(decodedData)

	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	urlSafeData, err := decodeURLSafeBase64(urlSafeBase64EncodedData)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(urlSafeData)
}