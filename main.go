package main

import (
	"bufio"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	a := app.New()
	w := a.NewWindow("gim")

	w.SetContent(widget.NewLabel(scanner.Text()))
	w.ShowAndRun()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
