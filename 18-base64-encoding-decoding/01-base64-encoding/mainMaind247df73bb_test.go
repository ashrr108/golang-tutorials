package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestMaind247df73bb(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		stdEncoded    string
		rawStdEncoded string
		urlEncoded    string
		rawUrlEncoded string
	}{
		{
			name:          "Basic String",
			input:         "Gol@ng is Awesome?~",
			stdEncoded:    "R29sQG5nIGlzIEF3ZXNvbWU/fw==",
			rawStdEncoded: "R29sQG5nIGlzIEF3ZXNvbWU/fw",
			urlEncoded:    "R29sQG5nIGlzIEF3ZXNvbWU_fw==",
			rawUrlEncoded: "R29sQG5nIGlzIEF3ZXNvbWU_fw",
		},
		{
			name:          "Empty String",
			input:         "",
			stdEncoded:    "",
			rawStdEncoded: "",
			urlEncoded:    "",
			rawUrlEncoded: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Standard Base64 Encoding
			encodedData := base64.StdEncoding.EncodeToString([]byte(tc.input))
			if encodedData != tc.stdEncoded {
				t.Errorf("Expected standard encoding: %s, got: %s", tc.stdEncoded, encodedData)
			}

			// Standard Base64 Encoding without padding
			encodedDataWithoutPadding := base64.RawStdEncoding.EncodeToString([]byte(tc.input))
			if encodedDataWithoutPadding != tc.rawStdEncoded {
				t.Errorf("Expected raw standard encoding: %s, got: %s", tc.rawStdEncoded, encodedDataWithoutPadding)
			}

			// URL and filename-safe Base64 encoding
			urlSafeEncodedData := base64.URLEncoding.EncodeToString([]byte(tc.input))
			if urlSafeEncodedData != tc.urlEncoded {
				t.Errorf("Expected URL safe encoding: %s, got: %s", tc.urlEncoded, urlSafeEncodedData)
			}

			// URL and filename-safe Base64 encoding without padding
			urlSafeEncodedDataWithoutPadding := base64.RawURLEncoding.EncodeToString([]byte(tc.input))
			if urlSafeEncodedDataWithoutPadding != tc.rawUrlEncoded {
				t.Errorf("Expected raw URL safe encoding: %s, got: %s", tc.rawUrlEncoded, urlSafeEncodedDataWithoutPadding)
			}
		})
	}
}

func main() {
	data := "Gol@ng is Awesome?~"

	// Standard Base64 Encoding
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedData)

	// Standard Base64 Encoding without padding
	encodedDataWithoutPadding := base64.RawStdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedDataWithoutPadding)

	// URL and filename-safe Base64 encoding
	urlSafeEncodedData := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlSafeEncodedData)

	// URL and filename-safe Base64 encoding without padding
	urlSafeEncodedDataWithoutPadding := base64.RawURLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlSafeEncodedDataWithoutPadding)
}