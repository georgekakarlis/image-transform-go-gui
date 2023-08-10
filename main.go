package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/chai2010/webp"
)

func main() {
	// Create a new application
	a := app.New()

	// Create a new window within the application
	w := a.NewWindow("Image Converter")

	// Create a new button widget labeled "Convert Image"
	// It takes a callback function which will be executed when the button is clicked
	convert := widget.NewButton("Convert Image", func() {

		// Open a file dialog, which allows the user to select a file
		// The callback function will be called once a file is selected or an error occurs
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {

			// If no error occurred and no file was selected, return early
			if err == nil && reader == nil {
				return
			}

			// If an error occurred, display it in a dialog and return early
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			// Ensure the reader gets closed once the function returns
			defer reader.Close()

			// Decode the image from the file
			img, _, err := image.Decode(reader)
			if err != nil {
				// If there's an error decoding the image, display it and return early
				dialog.ShowError(err, w)
				return
			}

			// Define output directory and output file path
			outputDir := "/Users/georgekakarlis/Desktop/goodprojects"
			outputFileName := getUniqueFileName(outputDir, "output", ".webp")

			// Create output file
			outputFile, err := os.Create(outputFileName)
			if err != nil {
				// If there's an error creating the output file, display it and return early
				dialog.ShowError(err, w)
				return
			}

			// Ensure the output file gets closed once the function returns
			defer outputFile.Close()

			// Define options for the WEBP encoding
			options := &webp.Options{
				Lossless: true,
				Quality:  90,
			}

			// Encode the image to WEBP and write it to the output file
			err = webp.Encode(outputFile, img, options)
			if err != nil {
				// If there's an error during the encoding, display it and return early
				dialog.ShowError(err, w)
				return
			}

			// Display a success message if the image was converted successfully
			dialog.ShowInformation("Success", "Image conversion successful!", w)
		}, w)

		// Set a filter to only show PNG, JPG and JPEG files in the dialog
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))

		// Show the file dialog
		fd.Show()
	})

	// Set the button as the content of the window
	w.SetContent(convert)

	// Show the window and run the application
	w.ShowAndRun()
}


func getUniqueFileName(baseDir, baseName, extension string) string {
	counter := 0
	filename := filepath.Join(baseDir, baseName+extension)
	// Keep checking and incrementing the counter until we find an unused filename
	for {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			return filename
		}
		counter++
		filename = filepath.Join(baseDir, fmt.Sprintf("%s_%d%s", baseName, counter, extension))
	}
}