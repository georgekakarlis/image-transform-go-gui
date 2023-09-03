package main

import (
	"fmt"
	"image"
	"os"

	"github.com/chai2010/webp"
)


func compressToWebP(inputFilePath, outputFilePath string, quality int) error {
	// Open the image file.
	file, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Decode the image.
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding image: %w", err)
	}
	if info, err := os.Stat(outputFilePath); err == nil && info.IsDir() {
		return fmt.Errorf("specified path is a directory, please provide a filename: %s", outputFilePath)
	}

	// Open a file for writing the compressed image.
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outFile.Close()

	// Encode the image to the .webp format with the specified quality.
	err = webp.Encode(outFile, img, &webp.Options{Quality: float32(quality)})
	if err != nil {
		return fmt.Errorf("error encoding to webp: %w", err)
	}

	return nil
}