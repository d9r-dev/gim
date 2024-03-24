package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("gim")

	w.SetContent(widget.NewLabel(string(file)))
	w.ShowAndRun()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
