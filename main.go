//go:generate fyne bundle -o bundled.go assets
package main

import (
	"fyne.io/fyne/v2/app"
)

const appID = "app.fylogin"

func main() {
	application := app.NewWithID(appID)

	application.SetIcon(resourceIconPng)

	NewLoginWindow(application).BuildAndShow()

	application.Run()
}
