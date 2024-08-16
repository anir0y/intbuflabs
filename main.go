package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:embed icon.png
var iconData []byte

func main() {
	// Create a new application
	myApp := app.New()
	myWindow := myApp.NewWindow("Arishti-Overflow")

	// Load the icon from embedded data
	icon := fyne.NewStaticResource("icon.png", iconData)
	myWindow.SetIcon(icon)

	// Create UI elements
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Phone Number")

	resultLabel := widget.NewLabel("")

	checkButton := widget.NewButton("Check Input", func() {
		if len(input.Text) < 10 {
			resultLabel.SetText("Invalid phone number.")
		} else if len(input.Text) == 10 {
			resultLabel.SetText("Valid phone number.")
		} else {
			resultLabel.SetText("Secret Code: Flag{Arishti-Overflow!}")
		}
	})

	// Layout
	content := container.NewVBox(
		input,
		checkButton,
		resultLabel,
	)

	// Set the content and display the window
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 169))
	myWindow.ShowAndRun()
}
