package main

import (
	"os"
	"testing"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func TestConvertToWebp(t *testing.T) {
	input := "/Users/georgekakarlis/Desktop/goodprojects/theinputimage.jpg"
	output := "/Users/georgekakarlis/Desktop/goodprojects/output.webp"
	quality := 40

	err := compressToWebP(input, output, quality)
	if err != nil {
		t.Fatalf("failed compressing to WebP: %v", err)
	}

	if !fileExists(output) {
		t.Fatalf("output file %s does not exist after compression", output)
	}
}
