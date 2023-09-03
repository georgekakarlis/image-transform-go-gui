package main

import (
	"time"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Compressor to .webp")

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Path to input image")

	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("Path to save compressed image")

	resultLabel := widget.NewLabel("")

	compressButton := widget.NewButton("Compress", func() {
		err := compressToWebP(inputEntry.Text, outputEntry.Text, 40)
		if err != nil {
			resultLabel.SetText("Error: " + err.Error())
		} else {
			resultLabel.SetText("Image compressed successfully!")
			go func() {
				time.Sleep(3 * time.Second)
				fyne.CurrentApp().SendNotification(fyne.NewNotification("Compression", "Resetting form..."))
				// Update UI on the main thread by refreshing the components
				inputEntry.SetText("")
				inputEntry.Refresh()  // Refresh the widget to reflect the change
				outputEntry.SetText("")
				outputEntry.Refresh() // Refresh the widget
				resultLabel.SetText("Form reset!")
				resultLabel.Refresh() // Refresh the widget
			}()
		}
	})

	 content := container.NewVBox(
		inputEntry,
		outputEntry,
		compressButton,
		resultLabel,
	) 

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

