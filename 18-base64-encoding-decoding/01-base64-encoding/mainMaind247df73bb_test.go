package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestMaind247df73bb(t *testing.T) {
	data := "Gol@ng is Awesome?~"

	// Standard Base64 Encoding
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	expectedEncodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="
	if encodedData != expectedEncodedData {
		t.Errorf("Expected %s but got %s", expectedEncodedData, encodedData)
	}

	// Standard Base64 Encoding without padding
	encodedDataWithoutPadding := base64.RawStdEncoding.EncodeToString([]byte(data))
	expectedEncodedDataWithoutPadding := "R29sQG5nIGlzIEF3ZXNvbWU_fg"
	if encodedDataWithoutPadding != expectedEncodedDataWithoutPadding {
		t.Errorf("Expected %s but got %s", expectedEncodedDataWithoutPadding, encodedDataWithoutPadding)
	}

	// URL and filename-safe Base64 encoding
	urlSafeEncodedData := base64.URLEncoding.EncodeToString([]byte(data))
	expectedUrlSafeEncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	if urlSafeEncodedData != expectedUrlSafeEncodedData {
		t.Errorf("Expected %s but got %s", expectedUrlSafeEncodedData, urlSafeEncodedData)
	}

	// URL and filename-safe Base64 encoding without padding
	urlSafeEncodedDataWithoutPadding := base64.RawURLEncoding.EncodeToString([]byte(data))
	expectedUrlSafeEncodedDataWithoutPadding := "R29sQG5nIGlzIEF3ZXNvbWU_fg"
	if urlSafeEncodedDataWithoutPadding != expectedUrlSafeEncodedDataWithoutPadding {
		t.Errorf("Expected %s but got %s", expectedUrlSafeEncodedDataWithoutPadding, urlSafeEncodedDataWithoutPadding)
	}
}

func TestMaind247df73bb_EmptyString(t *testing.T) {
	data := ""

	// Standard Base64 Encoding
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	expectedEncodedData := ""
	if encodedData != expectedEncodedData {
		t.Errorf("Expected %s but got %s", expectedEncodedData, encodedData)
	}

	// Standard Base64 Encoding without padding
	encodedDataWithoutPadding := base64.RawStdEncoding.EncodeToString([]byte(data))
	expectedEncodedDataWithoutPadding := ""
	if encodedDataWithoutPadding != expectedEncodedDataWithoutPadding {
		t.Errorf("Expected %s but got %s", expectedEncodedDataWithoutPadding, encodedDataWithoutPadding)
	}

	// URL and filename-safe Base64 encoding
	urlSafeEncodedData := base64.URLEncoding.EncodeToString([]byte(data))
	expectedUrlSafeEncodedData := ""
	if urlSafeEncodedData != expectedUrlSafeEncodedData {
		t.Errorf("Expected %s but got %s", expectedUrlSafeEncodedData, urlSafeEncodedData)
	}

	// URL and filename-safe Base64 encoding without padding
	urlSafeEncodedDataWithoutPadding := base64.RawURLEncoding.EncodeToString([]byte(data))
	expectedUrlSafeEncodedDataWithoutPadding := ""
	if urlSafeEncodedDataWithoutPadding != expectedUrlSafeEncodedDataWithoutPadding {
		t.Errorf("Expected %s but got %s", expectedUrlSafeEncodedDataWithoutPadding, urlSafeEncodedDataWithoutPadding)
	}
}