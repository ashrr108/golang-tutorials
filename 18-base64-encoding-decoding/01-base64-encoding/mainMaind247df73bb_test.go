package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func encodeData(data string) (string, string, string, string) {
	// Standard Base64 Encoding
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))

	// Standard Base64 Encoding without padding
	encodedDataWithoutPadding := base64.RawStdEncoding.EncodeToString([]byte(data))

	// URL and filename-safe Base64 encoding
	urlSafeEncodedData := base64.URLEncoding.EncodeToString([]byte(data))

	// URL and filename-safe Base64 encoding without padding
	urlSafeEncodedDataWithoutPadding := base64.RawURLEncoding.EncodeToString([]byte(data))

	return encodedData, encodedDataWithoutPadding, urlSafeEncodedData, urlSafeEncodedDataWithoutPadding
}

func TestEncodeData(t *testing.T) {
	data := "Gol@ng is Awesome?~"
	expectedEncodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
	expectedEncodedDataWithoutPadding := "R29sQG5nIGlzIEF3ZXNvbWU_fg"
	expectedUrlSafeEncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	expectedUrlSafeEncodedDataWithoutPadding := "R29sQG5nIGlzIEF3ZXNvbWU_fg"

	encodedData, encodedDataWithoutPadding, urlSafeEncodedData, urlSafeEncodedDataWithoutPadding := encodeData(data)

	if encodedData != expectedEncodedData {
		t.Errorf("Expected %s, but got %s", expectedEncodedData, encodedData)
	}

	if encodedDataWithoutPadding != expectedEncodedDataWithoutPadding {
		t.Errorf("Expected %s, but got %s", expectedEncodedDataWithoutPadding, encodedDataWithoutPadding)
	}

	if urlSafeEncodedData != expectedUrlSafeEncodedData {
		t.Errorf("Expected %s, but got %s", expectedUrlSafeEncodedData, urlSafeEncodedData)
	}

	if urlSafeEncodedDataWithoutPadding != expectedUrlSafeEncodedDataWithoutPadding {
		t.Errorf("Expected %s, but got %s", expectedUrlSafeEncodedDataWithoutPadding, urlSafeEncodedDataWithoutPadding)
	}
}

func TestEncodeDataWithEmptyString(t *testing.T) {
	data := ""
	expectedEncodedData := ""
	expectedEncodedDataWithoutPadding := ""
	expectedUrlSafeEncodedData := ""
	expectedUrlSafeEncodedDataWithoutPadding := ""

	encodedData, encodedDataWithoutPadding, urlSafeEncodedData, urlSafeEncodedDataWithoutPadding := encodeData(data)

	if encodedData != expectedEncodedData {
		t.Errorf("Expected empty string, but got %s", encodedData)
	}

	if encodedDataWithoutPadding != expectedEncodedDataWithoutPadding {
		t.Errorf("Expected empty string, but got %s", encodedDataWithoutPadding)
	}

	if urlSafeEncodedData != expectedUrlSafeEncodedData {
		t.Errorf("Expected empty string, but got %s", urlSafeEncodedData)
	}

	if urlSafeEncodedDataWithoutPadding != expectedUrlSafeEncodedDataWithoutPadding {
		t.Errorf("Expected empty string, but got %s", urlSafeEncodedDataWithoutPadding)
	}
}